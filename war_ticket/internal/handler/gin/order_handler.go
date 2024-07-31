package gin_handler

import (
	"github.com/benebobaa/valo"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/usecase"
	"war_ticket/pkg"
)

type OrderHandlerImpl struct {
	orderUsecase usecase.OrderUsecase
}

type OrderHandler interface {
	interfaces.GinHandler
}

func NewOrderHandler(oc usecase.OrderUsecase) OrderHandler {
	return &OrderHandlerImpl{
		orderUsecase: oc,
	}
}

// Create implements OrderHandler.
func (o *OrderHandlerImpl) Create(c *gin.Context) {
	var logger pkg.LogFormat
	var request dto.OrderRequest

	startTime := time.Now()

	defer func() {
		logger.ProcessTime = uint(time.Since(startTime).Milliseconds())
		logger.Execute()
	}()

	if err := c.ShouldBindJSON(&request); err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Failed to parse request body",
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, dto.BaseResponse[*domain.Event]{
			Error: err.Error(),
		})
		return
	}

	if err := valo.Validate(request); err != nil {
		logger = pkg.LogFormat{
			HttpStatus: http.StatusBadRequest,
			Message:    "Error validating request",
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, dto.BaseResponse[*domain.Event]{
			Error: err.Error(),
		})
		return
	}

	result, err := o.orderUsecase.CreateOrder(c.Request.Context(), &request)

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Failed to create order",
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, dto.BaseResponse[*domain.Event]{
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

	c.JSON(http.StatusCreated, dto.BaseResponse[*domain.Order]{
		Data: result,
	})
}

// FindAll implements OrderHandler.
func (o *OrderHandlerImpl) FindAll(c *gin.Context) {
	var logger pkg.LogFormat

	startTime := time.Now()

	defer func() {
		logger.ProcessTime = uint(time.Since(startTime).Milliseconds())
		logger.Execute()
	}()

	result := o.orderUsecase.GetAll()

	logger = pkg.LogFormat{
		IsSuccess:  true,
		Message:    "Successfully retrieved all orders",
		HttpStatus: http.StatusOK,
	}

	c.JSON(http.StatusOK, dto.BaseResponse[[]domain.Order]{
		Data: result,
	})
}
