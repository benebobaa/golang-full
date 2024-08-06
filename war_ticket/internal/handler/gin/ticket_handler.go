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

type TicketHandlerImpl struct {
	ticketUsecase usecase.TicketUsecase
}

type TicketHandler interface {
	interfaces.GinHandler
}

func NewTicketHandler(tc usecase.TicketUsecase) TicketHandler {
	return &TicketHandlerImpl{
		ticketUsecase: tc,
	}
}

// Create implements TicketHandler.
func (t *TicketHandlerImpl) Create(c *gin.Context) {
	var request dto.TicketRequest
	var logger pkg.LogFormat

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
		c.JSON(http.StatusBadRequest, dto.BaseResponse[*dto.TicketResponse]{
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

	result, err := t.ticketUsecase.Save(c.Request.Context(), &request)

	if err != nil {
		logger = pkg.LogFormat{
			IsSuccess:  false,
			HttpStatus: http.StatusBadRequest,
			Message:    "Failed to save ticket event",
			Error:      err.Error(),
		}
		c.JSON(http.StatusBadRequest, dto.BaseResponse[*dto.TicketResponse]{
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

	c.JSON(http.StatusCreated, dto.BaseResponse[*dto.TicketResponse]{
		Data: result,
	})
}

// FindAll implements TicketHandler.
func (t *TicketHandlerImpl) FindAll(c *gin.Context) {
	var logger pkg.LogFormat
	startTime := time.Now()
	defer func() {
		logger.ProcessTime = uint(time.Since(startTime).Milliseconds())
		logger.Execute()
	}()

	query := c.Query("include")
	var result any
	var err error

	if query == "event" {
		result, err = t.ticketUsecase.GetAllWithEvent()
		if err != nil {
			logger = pkg.LogFormat{
				IsSuccess:  false,
				HttpStatus: http.StatusInternalServerError,
				Message:    "Failed to get tickets with event",
				Error:      err.Error(),
			}
			c.JSON(http.StatusInternalServerError, dto.BaseResponse[any]{
				Error: err.Error(),
			})
			return
		}
		logger.Message = "Successfully got all tickets with event"
	} else {
		result = t.ticketUsecase.GetAll()
		logger.Message = "Successfully got all tickets"
	}

	logger.IsSuccess = true
	logger.HttpStatus = http.StatusOK
	logger.Data = result

	c.JSON(http.StatusOK, dto.BaseResponse[any]{
		Data: result,
	})
}
