package main

import (
	"war_ticket/internal/domain"
	"war_ticket/internal/handler"
	"war_ticket/internal/repository"
	"war_ticket/internal/usecase"
	"war_ticket/pkg"
)

func initHandler() (handler.EventHandler, handler.TicketHandler, handler.OrderHandler) {

	er := repository.NewEventRepository()
	ec := usecase.NewEventUsecase(er)
	eh := handler.NewEventHandler(ec)

	ter := repository.NewTicketEventRepository()

	tr := repository.NewTicketRepository()
	tc := usecase.NewTicketUsecase(er, tr, ter)
	th := handler.NewTicketHandler(tc)

	or := repository.NewOrderRepository()
	oc := usecase.NewOrderUsecase(or, tr)
	oh := handler.NewOrderHandler(oc)

	generateEvent(ec)

	return eh, th, oh
}

func initRouter(
	eventHandler handler.EventHandler,
	ticketHandler handler.TicketHandler,
	orderHandler handler.OrderHandler,
) *pkg.Router {

	router := pkg.NewRouter()

	// event
	router.GET("/api/events", eventHandler.FindAll)
	router.POST("/api/events", eventHandler.Create)

	// ticket
	router.GET("/api/tickets", ticketHandler.FindAll)
	router.POST("/api/tickets", ticketHandler.Create)

	// order
	router.GET("/api/orders", orderHandler.FindAll)
	router.POST("/api/orders", orderHandler.Create)
	router.GET("/api/orders/length", orderHandler.GetTotalElements)

	return router
}

func generateEvent(ec usecase.EventUsecase) {
	event1 := domain.Event{
		Name:     "Lomba joget",
		Location: "Jaksel",
	}

	event2 := domain.Event{
		Name:     "Konser Nyanyi",
		Location: "Blok M",
	}

	ec.Save(&event1)
	ec.Save(&event2)
}
