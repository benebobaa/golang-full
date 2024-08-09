package app

import (
	"orchestra-svc/internal/delivery/http"
	"orchestra-svc/internal/delivery/kafka"
	"orchestra-svc/internal/repository/sqlc"
	"orchestra-svc/internal/usecase"
)

func (app *App) startService() error {

	userProducer, err := kafka.NewKafkaProducer(
		[]string{app.config.KafkaBroker},
		app.config.UserTopic,
	)

	if err != nil {
		return err
	}

	app.msg = kafka.NewMessageHandler(userProducer)

	sqlc := sqlc.New(app.db)

	workflowUsecase := usecase.NewWorkflowUsecase(sqlc)
	workflowHandler := http.NewWorkflowHandler(workflowUsecase)

	wfGroupv1 := app.gin.Group("/v1/workflow")
	workflowHandler.RegisterRoutes(wfGroupv1)

	return nil
}
