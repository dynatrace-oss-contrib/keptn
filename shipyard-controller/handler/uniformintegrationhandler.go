package handler

import (
	"github.com/gin-gonic/gin"
	keptnmodels "github.com/keptn/go-utils/pkg/api/models"
	"github.com/keptn/keptn/shipyard-controller/models"
	"net/http"
	"time"
)

type IUniformIntegrationHandler interface {
	Register(context *gin.Context)
	KeepAlive(context *gin.Context)
	Unregister(context *gin.Context)
	GetRegistrations(context *gin.Context)
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

	// for backwards compatibility, we check if there is a Subscriptions field set
	// if not, we are taking the old Subscription field and map it to the new Subscriptions field
	if integration.Subscriptions == nil {
		ts := keptnmodels.TopicSubscription{
			Topics: integration.Subscription.Topics,
			Status: integration.Subscription.Status,
			Filter: keptnmodels.TopicSubscriptionFilter{
				Projects: []string{integration.Subscription.Filter.Project},
				Stages:   []string{integration.Subscription.Filter.Stage},
				Services: []string{integration.Subscription.Filter.Service},
			},
		}
		integration.Subscriptions = append(integration.Subscriptions, ts)
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

	integration.ID = hash

	integration.MetaData.LastSeen = time.Now().UTC()

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
// @Param id path string true "id"
// @Success 200
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Router /uniform/registration/{id} [delete]
func (rh *UniformIntegrationHandler) Unregister(c *gin.Context) {
	integrationID := c.Param("id")

	if err := rh.integrationManager.Unregister(integrationID); err != nil {
		SetInternalServerErrorResponse(err, c)
		return
	}
	c.JSON(http.StatusOK, &models.UnregisterResponse{})
}

// GetRegistrations Retrieves uniform integrations matching the provided filter
// @Summary Retrieve uniform integrations
// @Description Retrieve uniform integrations
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

// KeepAlive updates a uniform integration and returns its current registration data
// @Summary Updates a uniform integration and returns its current registration data
// @Description Updates a uniform integration and returns its current registration data
// @Tags Uniform
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Integration "ok"
// @Failure 400 {object} models.Error "Invalid payload"
// @Failure 500 {object} models.Error "Internal error"
// @Router /uniform/registration/{id} [put]
func (rh *UniformIntegrationHandler) KeepAlive(c *gin.Context) {
	integrationID := c.Param("id")
	registrations, err := rh.integrationManager.GetRegistrations(models.GetUniformIntegrationsParams{ID: integrationID})
	if err != nil {
		SetInternalServerErrorResponse(err, c)
		return
	}

	if len(registrations) != 1 {
		SetNotFoundErrorResponse(nil, c, "No registration with id "+integrationID+"found")
		return
	}
	registration := registrations[0]
	registration.MetaData.LastSeen = time.Now().UTC()

	if err := rh.integrationManager.Register(registration); err != nil {
		SetInternalServerErrorResponse(err, c)
		return
	}

	c.JSON(http.StatusOK, registration)

}
