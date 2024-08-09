package app

import (
	"time"
	"user-svc/internal/usecase"
	"user-svc/pkg/http_client"

	"github.com/gin-gonic/gin"
)

func (app *App) startService() error {

	userClient := http_client.NewUserClient(
		"https://beneboba.wiremockapi.cloud/users/validate",
		5*time.Second,
	)

	usecase := usecase.NewUsecase(userClient)
	app.usecase = usecase

	app.gin.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	return nil
}
