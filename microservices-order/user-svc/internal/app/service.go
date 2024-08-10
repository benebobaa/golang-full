package app

import (
	"time"
	"user-svc/internal/delivery/messaging"
	"user-svc/internal/usecase"
	"user-svc/pkg/http_client"
	"user-svc/pkg/producer"

	"github.com/gin-gonic/gin"
)

func (app *App) startService(orchestraProducer *producer.KafkaProducer) error {

	userClient := http_client.NewUserClient(
		app.config.ClientUrl,
		5*time.Second,
	)

	usecase := usecase.NewUsecase(userClient, orchestraProducer)

	app.msg = messaging.NewMessageHandler(usecase)

	app.gin.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	return nil
}
