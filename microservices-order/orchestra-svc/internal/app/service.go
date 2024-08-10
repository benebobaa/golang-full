package app

import (
	"orchestra-svc/internal/delivery/http"
	"orchestra-svc/internal/delivery/messaging"
	"orchestra-svc/internal/repository/sqlc"
	"orchestra-svc/internal/usecase"
	"orchestra-svc/pkg/producer"
)

func (app *App) startService(producer *producer.KafkaProducer) error {

	oc := usecase.NewOrderUsecase(producer)

	app.msg = messaging.NewMessageHandler(oc)

	sqlc := sqlc.New(app.db)

	wfu := usecase.NewWorkflowUsecase(sqlc)
	wfh := http.NewWorkflowHandler(wfu)

	wfGroupv1 := app.gin.Group("/api/v1/workflow")
	wfh.RegisterRoutes(wfGroupv1)

	return nil
}
