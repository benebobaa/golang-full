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

type TicketHandlerImpl struct {
	ticketUsecase usecase.TicketUsecase
}

type TicketHandler interface {
	interfaces.Handler
}

func NewTicketHandler(tc usecase.TicketUsecase) TicketHandler {
	return &TicketHandlerImpl{
		ticketUsecase: tc,
	}
}

// Create implements TicketHandler.
func (t *TicketHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var request dto.TicketRequest
	var logger pkg.LogFormat

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
			dto.BaseResponse[*dto.TicketResponse]{
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

	result, err := t.ticketUsecase.Save(r.Context(), &request)

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Failed save ticket event",
			Error:      err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.WriteToResponseBody(
			w,
			dto.BaseResponse[*dto.TicketResponse]{
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

	w.WriteHeader(http.StatusOK)
	json.WriteToResponseBody(
		w,
		dto.BaseResponse[*dto.TicketResponse]{
			Data: result,
		})

}

// FindAll implements TicketHandler.
func (t *TicketHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	var logger pkg.LogFormat
	startTime := time.Now()
	defer func() {
		logger.ProcessTime = uint(time.Since(startTime).Milliseconds())
		logger.Execute()
	}()

	query := r.URL.Query().Get("include")
	var result any
	var err error

	if query == "event" {
		result, err = t.ticketUsecase.GetAllWithEvent()
		if err != nil {
			logger.IsSuccess = false
			logger.HttpStatus = http.StatusInternalServerError
			logger.Message = "Failed to get tickets with event"
			logger.Error = err.Error()

			w.WriteHeader(http.StatusInternalServerError)
			json.WriteToResponseBody(w, dto.BaseResponse[any]{
				Error: err.Error(),
			})
			return
		}
		logger.Message = "Success get all tickets with event"
	} else {
		result = t.ticketUsecase.GetAll()
		logger.Message = "Success get all tickets"
	}

	logger.IsSuccess = true
	logger.HttpStatus = http.StatusOK
	logger.Data = result

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.WriteToResponseBody(w, dto.BaseResponse[any]{
		Data: result,
	})
}
