package handler

import (
	"net/http"
	"time"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/json"
	"war_ticket/internal/middleware"
	"war_ticket/internal/usecase"
	"war_ticket/pkg"

	"github.com/benebobaa/valo"
)

type OrderHandlerImpl struct {
	orderUsecase usecase.OrderUsecase
}

type OrderHandler interface {
	interfaces.Handler
	GetTotalElements(w http.ResponseWriter, r *http.Request)
}

func NewOrderHandler(oc usecase.OrderUsecase) OrderHandler {
	return &OrderHandlerImpl{
		orderUsecase: oc,
	}
}

// Create implements OrderHandler.
func (o *OrderHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var logger pkg.LogFormat
	var request dto.OrderRequest

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

	user := r.Context().Value(middleware.ContextUserKey).(*domain.User)

	if user == nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusUnauthorized,
			Message:    "Unauthorized request",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.WriteToResponseBody(
			w,
			dto.BaseResponse[*domain.Event]{
				Error: "Unauthorized request",
			})
		return
	}

	result, err := o.orderUsecase.CreateOrder(r.Context(), &request, user)

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Failed to create order",
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

	logger = pkg.LogFormat{
		IsSuccess:  true,
		HttpStatus: http.StatusCreated,
		Message:    "Successfully created order",
		Data:       result,
	}

	w.WriteHeader(http.StatusCreated)
	json.WriteToResponseBody(
		w,
		dto.BaseResponse[*domain.Order]{
			Data: result,
		})
}

// FindAll implements OrderHandler.
func (o *OrderHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	var logger pkg.LogFormat

	startTime := time.Now()

	defer func() {
		logger.ProcessTime = uint(time.Now().Sub(startTime).Milliseconds())
		logger.Execute()
	}()

	result := o.orderUsecase.GetAll()

	w.Header().Add("Content-Type", "application/json")

	logger = pkg.LogFormat{
		IsSuccess:  true,
		HttpStatus: http.StatusOK,
		Message:    "Successfully retrieved all orders",
		Data:       result,
	}

	w.WriteHeader(http.StatusOK)
	json.WriteToResponseBody(
		w,
		dto.BaseResponse[[]domain.Order]{
			Data: result,
		})
}

// GetTotalElements implements OrderHandler.
func (o *OrderHandlerImpl) GetTotalElements(w http.ResponseWriter, r *http.Request) {

	result := o.orderUsecase.GetAll()

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.WriteToResponseBody(
		w,
		dto.BaseResponse[int]{
			Data: len(result),
		})
}
