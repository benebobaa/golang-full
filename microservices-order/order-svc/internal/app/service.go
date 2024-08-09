package app

import (
	"order-svc/internal/delivery/http"
	"order-svc/internal/delivery/kafka"
	"order-svc/internal/middleware"
	"order-svc/internal/repository/sqlc"
	"order-svc/internal/usecase"
)

func (app *App) startService() error {

	sqlc := sqlc.New(app.db)

	orchestraProducer, err := kafka.NewKafkaProducer(
		[]string{app.config.KafkaBroker},
		app.config.OrchestraTopic,
	)

	app.msg = kafka.NewMessageHandler(orchestraProducer)

	if err != nil {
		return err
	}

	orderUsecase := usecase.NewOrderUsecase(sqlc, orchestraProducer)

	authHandler := http.NewAuthHandler()
	orderHandler := http.NewOrderHandler(orderUsecase)

	apiV1 := app.gin.Group("/api/v1")
	authV1 := apiV1.Group("/auth")
	orderV1 := apiV1.Group("/order")

	orderV1.Use(middleware.AuthMiddleware())

	authHandler.RegisterRoutes(authV1)
	orderHandler.RegisterRoutes(orderV1)

	return nil
}
