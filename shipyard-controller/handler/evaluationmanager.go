package handler

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/keptn/go-utils/pkg/common/strutils"
	"github.com/keptn/go-utils/pkg/common/timeutils"
	"github.com/keptn/go-utils/pkg/lib/keptn"
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
	"github.com/keptn/keptn/shipyard-controller/common"
	"github.com/keptn/keptn/shipyard-controller/db"
	"github.com/keptn/keptn/shipyard-controller/models"
	"github.com/keptn/keptn/shipyard-controller/operations"
)

const userFriendlyTimeFormat = "2006-01-02T15:04:05"

const (
	evaluationErrInvalidTimeframe = iota
	evaluationErrSendEventFailed
	evaluationErrServiceNotAvailable
)

//go:generate moq -pkg fake -skip-ensure -out ./fake/evaluationmanager.go . IEvaluationManager
type IEvaluationManager interface {
	CreateEvaluation(ctx context.Context, project, stage, service string, params *operations.CreateEvaluationParams) (*operations.CreateEvaluationResponse, *models.Error)
}

type EvaluationManager struct {
	EventSender keptn.EventSender
	ServiceAPI  db.ServicesDbOperations
}

func NewEvaluationManager(eventSender keptn.EventSender, serviceAPI db.ServicesDbOperations) (*EvaluationManager, error) {
	return &EvaluationManager{
		EventSender: eventSender,
		ServiceAPI:  serviceAPI,
	}, nil

}

func (em *EvaluationManager) CreateEvaluation(ctx context.Context, project, stage, service string, params *operations.CreateEvaluationParams) (*operations.CreateEvaluationResponse, *models.Error) {
	_, err := em.ServiceAPI.GetService(project, stage, service)
	if err != nil {
		return nil, &models.Error{
			Code:    evaluationErrServiceNotAvailable,
			Message: strutils.Stringp(err.Error()),
		}
	}

	keptnContext := uuid.New().String()
	extensions := make(map[string]interface{})
	extensions["shkeptncontext"] = keptnContext

	var start, end *time.Time
	start, end, err = timeutils.GetStartEndTime(timeutils.GetStartEndTimeParams{
		StartDate: params.Start,
		EndDate:   params.End,
		Timeframe: params.Timeframe,
	})
	if err != nil {
		// if we got an error, try again with other time format
		start, end, err = timeutils.GetStartEndTime(timeutils.GetStartEndTimeParams{
			StartDate:  params.Start,
			EndDate:    params.End,
			Timeframe:  params.Timeframe,
			TimeFormat: userFriendlyTimeFormat,
		})
		if err != nil {
			return nil, &models.Error{
				Code:    evaluationErrInvalidTimeframe,
				Message: strutils.Stringp(err.Error()),
			}
		}
	}

	eventContext := &operations.CreateEvaluationResponse{KeptnContext: keptnContext}

	evaluationTriggeredEvent := keptnv2.EvaluationTriggeredEventData{
		EventData: keptnv2.EventData{
			Project: project,
			Service: service,
			Stage:   stage,
			Labels:  params.Labels,
		},
		Evaluation: keptnv2.Evaluation{
			Start: timeutils.GetKeptnTimeStamp(*start),
			End:   timeutils.GetKeptnTimeStamp(*end),
		},
	}

	ce := common.CreateEventWithPayload(keptnContext, "", keptnv2.GetTriggeredEventType(stage+"."+keptnv2.EvaluationTaskName), evaluationTriggeredEvent)
	if err := ce.Context.SetSource("https://github.com/keptn/keptn/api"); err != nil {
		return nil, &models.Error{
			Code:    evaluationErrSendEventFailed,
			Message: common.Stringp(err.Error()),
		}
	}

	// TODO: Should we add a target property also to type keptn.EventSender to avoid the conversion here?
	var target string
	if httpEventSender, ok := em.EventSender.(*keptnv2.HTTPEventSender); ok {
		target = httpEventSender.EventsEndpoint
	}

	ctx = cloudevents.ContextWithTarget(ctx, target)
	ctx = cloudevents.WithEncodingStructured(ctx)

	if err := em.EventSender.Send(ctx, ce); err != nil {
		return nil, &models.Error{
			Code:    evaluationErrSendEventFailed,
			Message: common.Stringp(err.Error()),
		}
	}

	return eventContext, nil
}
