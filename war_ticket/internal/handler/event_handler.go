package handler

import (
	"net/http"
	"time"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/json"
	"war_ticket/internal/usecase"
	"war_ticket/pkg"

	"github.com/benebobaa/valo"
)

type EventHandlerImpl struct {
	eventUsecase usecase.EventUsecase
}

type EventHandler interface {
	interfaces.Handler
}

func NewEventHandler(eventUsecase usecase.EventUsecase) EventHandler {
	return &EventHandlerImpl{
		eventUsecase: eventUsecase,
	}
}

// Create implements EventHandler.
func (e *EventHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var logger pkg.LogFormat
	var request domain.Event

	startTime := time.Now()

	defer func() {
		logger.ProcessTime = uint(time.Now().Sub(startTime).Milliseconds())
		logger.Execute()
	}()

	err := json.ReadFromRequestBody(r, &request)
	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Failed parse read from request body",
			Error:      err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.WriteToResponseBody(
			w,
			dto.BaseResponse[*domain.Event]{
				Error: err.Error(),
			})
		return
	}

	err = valo.Validate(request)

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Error validation request",
			Error:      err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.WriteToResponseBody(
			w,
			dto.BaseResponse[*domain.Event]{
				Error: err.Error(),
			})
		return
	}

	result, err := e.eventUsecase.Save(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.WriteToResponseBody(
			w,
			dto.BaseResponse[*domain.Event]{
				Error: err.Error(),
			})
		return
	}

	logger = pkg.LogFormat{
		IsSuccess:  true,
		HttpStatus: http.StatusCreated,
		Message:    "Success created event",
		Data:       result,
	}

	w.WriteHeader(http.StatusCreated)
	json.WriteToResponseBody(
		w,
		dto.BaseResponse[*domain.Event]{
			Data: result,
		})
}

// FindAll implements EventHandler.
func (e *EventHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	result := e.eventUsecase.GetAll()

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.WriteToResponseBody(
		w,
		dto.BaseResponse[[]domain.Event]{
			Data: result,
		})
}
