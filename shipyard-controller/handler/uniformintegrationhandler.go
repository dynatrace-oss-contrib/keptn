package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	keptnmodels "github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/keptn/shipyard-controller/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type IUniformIntegrationHandler interface {
	Register(context *gin.Context)
	KeepAlive(context *gin.Context)
	Unregister(context *gin.Context)
	GetRegistrations(context *gin.Context)
	GetSubscription(context *gin.Context)
	GetSubscriptions(c *gin.Context)
	CreateSubscription(c *gin.Context)
	DeleteSubscription(c *gin.Context)
	UpdateSubscription(c *gin.Context)
}

type UniformIntegrationHandler struct {
	integrationManager IUniformIntegrationManager
}

func NewUniformIntegrationHandler(im IUniformIntegrationManager) *UniformIntegrationHandler {
	return &UniformIntegrationHandler{integrationManager: im}
}

// Register creates or updates a uniform integration
// @Summary Register a uniform integration
// @Description Register a uniform integration
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integration body models.Integration true "Integration"
// @Success 200
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Router /uniform/registration [post]
func (rh *UniformIntegrationHandler) Register(c *gin.Context) {
	integration := &models.Integration{}

	if err := c.ShouldBindJSON(integration); err != nil {
		SetBadRequestErrorResponse(err, c)
		return
	}

	integrationID := keptnmodels.IntegrationID{
		Name:      integration.Name,
		Namespace: integration.MetaData.KubernetesMetaData.Namespace,
		NodeName:  integration.MetaData.Hostname,
	}

	hash, err := integrationID.Hash()
	if err != nil {
		SetBadRequestErrorResponse(err, c)
		return
	}

	//setting IDs and last seen timestamp
	integration.ID = hash
	integration.MetaData.LastSeen = time.Now().UTC()
	for i := range integration.Subscriptions {
		s := &integration.Subscriptions[i]
		s.ID = uuid.New().String()
	}

	// for backwards compatibility, we check if there is a Subscriptions field set
	// if not, we are taking the old Subscription field and map it to the new Subscriptions field
	// Note: "old" registrations will NOT get subscription IDs
	if integration.Subscriptions == nil {
		topic := ""
		if len(integration.Subscription.Topics) > 0 {
			topic = integration.Subscription.Topics[0]
		}
		ts := keptnmodels.EventSubscription{
			Event: topic,
			Filter: keptnmodels.EventSubscriptionFilter{
				Projects: []string{integration.Subscription.Filter.Project},
				Stages:   []string{integration.Subscription.Filter.Stage},
				Services: []string{integration.Subscription.Filter.Service},
			},
		}
		integration.Subscriptions = append(integration.Subscriptions, ts)
	}

	if err := rh.integrationManager.Register(*integration); err != nil {
		SetInternalServerErrorResponse(err, c)
		return
	}
	c.JSON(http.StatusOK, &models.RegisterResponse{
		ID: integration.ID,
	})
}

// Unregister deletes a uniform integration
// @Summary Unregister a uniform integration
// @Description Unregister a uniform integration
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "integrationID"
// @Success 200
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Router /uniform/registration/{integrationID} [delete]
func (rh *UniformIntegrationHandler) Unregister(c *gin.Context) {
	integrationID := c.Param("integrationID")

	if err := rh.integrationManager.Unregister(integrationID); err != nil {
		SetInternalServerErrorResponse(err, c)
		return
	}
	c.JSON(http.StatusOK, &models.UnregisterResponse{})
}

// GetRegistrations Retrieve uniform integrations matching the provided filter
// @Summary Retrieve uniform integrations matching the provided filter
// @Description Retrieve uniform integrations matching the provided filter
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Param name query string false "name"
// @Param project query string false "project"
// @Param stage query string false "stage"
// @Param service query string false "service"
// @Success 200 {object} []models.Integration "ok"
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Router /uniform/registration [get]
func (rh *UniformIntegrationHandler) GetRegistrations(c *gin.Context) {
	params := &models.GetUniformIntegrationsParams{}
	if err := c.ShouldBindQuery(params); err != nil {
		SetBadRequestErrorResponse(err, c, "Invalid request format")
		return
	}
	uniformIntegrations, err := rh.integrationManager.GetRegistrations(*params)
	if err != nil {
		SetInternalServerErrorResponse(err, c, "Unable to query uniform integrations repository")
		return
	}

	c.JSON(http.StatusOK, uniformIntegrations)
}

// KeepAlive returns current registration data of an integration
// @Summary Heartbeat for uniform integrations
// @Description Heartbeat for uniform integrations
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integrationID path string true "integrationID"
// @Success 200 {object} models.Integration "ok"
// @Failure 404 {object} models.Error "Not found"
// @Failure 500 {object} models.Error "Internal error"
// @Router /uniform/registration/{integrationID}/ping [PUT]
func (rh *UniformIntegrationHandler) KeepAlive(c *gin.Context) {
	integrationID := c.Param("integrationID")

	registration, err := rh.integrationManager.UpdateLastSeen(integrationID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			SetNotFoundErrorResponse(err, c)
			return
		}
		SetInternalServerErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, registration)

}

// CreateSubscription creates a new subscription
// @Summary  Create a new subscription
// @Description  Create a new subscription
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integrationID path string true "integrationID"
// @Param subscription body models.Subscription true "Subscription"
// @Success 201
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Failure 404 {object} models.Error "Not found"
// @Router /uniform/registration/{integrationID}/subscription [post]
func (rh *UniformIntegrationHandler) CreateSubscription(c *gin.Context) {

	integrationID := c.Param("integrationID")
	subscription := &models.Subscription{}

	if err := c.ShouldBindJSON(subscription); err != nil {
		SetBadRequestErrorResponse(err, c)
		return
	}
	subscription.ID = uuid.New().String()

	err := rh.integrationManager.CreateOrUpdateSubscription(integrationID, *subscription)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			SetNotFoundErrorResponse(err, c)
			return
		}
		SetInternalServerErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, &models.CreateSubscriptionResponse{
		ID: subscription.ID,
	})
}

// UpdateSubscription updates or creates a subscription
// @Summary  Update or create a subscription
// @Description Update or create a subscription
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integrationID path string true "integrationID"
// @Param subscriptionID path string true "subscriptionID"
// @Param subscription body models.Subscription true "Subscription"
// @Success 201
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Failure 404 {object} models.Error "Not found"
// @Router /uniform/registration/{integrationID}/subscription/{subscriptionID} [put]
func (rh *UniformIntegrationHandler) UpdateSubscription(c *gin.Context) {

	integrationID := c.Param("integrationID")
	subscriptionID := c.Param("subscriptionID")

	subscription := &models.Subscription{}

	if err := c.ShouldBindJSON(subscription); err != nil {
		SetBadRequestErrorResponse(err, c)
		return
	}
	subscription.ID = subscriptionID

	err := rh.integrationManager.CreateOrUpdateSubscription(integrationID, *subscription)
	if err != nil {
		//TODO: set appropriate http codes
		SetInternalServerErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusCreated, &models.CreateSubscriptionResponse{
		ID: subscription.ID,
	})
}

// DeleteSubscription deletes a new subscription
// @Summary  Delete a subscription
// @Description  Delete a subscription
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integrationID path string true "integrationID"
// @Param subscriptionID path string true "subscriptionID"
// @Success 200
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Failure 404 {object} models.Error "Not found"
// @Router /uniform/registration/{integrationID}/subscription/{subscriptionID} [delete]
func (rh *UniformIntegrationHandler) DeleteSubscription(c *gin.Context) {
	integrationID := c.Param("integrationID")
	subscriptionID := c.Param("subscriptionID")

	err := rh.integrationManager.DeleteSubscription(integrationID, subscriptionID)
	if err != nil {
		SetInternalServerErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, models.DeleteSubscriptionResponse{})
}

// GetSubscription retrieves an already existing subscription
// @Summary  Retrieve an already existing subscription
// @Description  Retrieve an already existing subscription
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integrationID path string true "integrationID"
// @Param subscriptionID path string true "subscriptionID"
// @Success 200 {object} models.Subscription "ok"
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Failure 404 {object} models.Error "Not found"
// @Router /uniform/registration/{integrationID}/subscription/{subscriptionID} [get]
func (rh *UniformIntegrationHandler) GetSubscription(c *gin.Context) {
	integrationID := c.Param("integrationID")
	subscriptionID := c.Param("subscriptionID")

	subscription, err := rh.integrationManager.GetSubscription(integrationID, subscriptionID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			SetNotFoundErrorResponse(err, c)
			return
		}
		SetInternalServerErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, subscription)
}

// GetSubscriptions retrieves all subscriptions of a uniform integration
// @Summary  Retrieve all subscriptions of a uniform integration
// @Description  Retrieve all subscriptions of a uniform integration
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param integrationID path string true "integrationID"
// @Success 200 {object} []models.Subscription "ok"
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Failure 404 {object} models.Error "Not found"
// @Router /uniform/registration/{integrationID}/subscription [get]
func (rh *UniformIntegrationHandler) GetSubscriptions(c *gin.Context) {
	integrationID := c.Param("integrationID")

	subscriptions, err := rh.integrationManager.GetSubscriptions(integrationID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			SetNotFoundErrorResponse(err, c)
			return
		}
		SetInternalServerErrorResponse(err, c)
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}
