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

type EventHandlerImpl struct {
	eventUsecase usecase.EventUsecase
}

type EventHandler interface {
	interfaces.GinHandler
}

func NewEventHandler(eventUsecase usecase.EventUsecase) EventHandler {
	return &EventHandlerImpl{
		eventUsecase: eventUsecase,
	}
}

// Create implements EventHandler.
func (e *EventHandlerImpl) Create(c *gin.Context) {
	var logger pkg.LogFormat
	var request domain.Event

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
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Error validating request",
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, dto.BaseResponse[*domain.Event]{
			Error: err.Error(),
		})
		return
	}

	result, err := e.eventUsecase.Save(c.Request.Context(), &request)

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusInternalServerError,
			Message:    "Error creating event",
			Error:      err.Error(),
		}
		c.JSON(http.StatusInternalServerError, dto.BaseResponse[*domain.Event]{
			Error: err.Error(),
		})
		return
	}

	logger = pkg.LogFormat{
		IsSuccess:  true,
		HttpStatus: http.StatusCreated,
		Message:    "Successfully created event",
		Data:       result,
	}

	c.JSON(http.StatusCreated, dto.BaseResponse[*domain.Event]{
		Data: result,
	})
}

// FindAll implements EventHandler.
func (e *EventHandlerImpl) FindAll(c *gin.Context) {
	result := e.eventUsecase.GetAll()

	c.JSON(http.StatusOK, dto.BaseResponse[[]domain.Event]{
		Data: result,
	})
}
