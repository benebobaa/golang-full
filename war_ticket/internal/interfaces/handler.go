package interfaces

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type GinHandler interface {
	FindAll(ctx *gin.Context)
	Create(ctx *gin.Context)
}
