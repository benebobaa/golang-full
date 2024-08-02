package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"war_ticket/internal/handler"
	ginhandler "war_ticket/internal/handler/gin"
	"war_ticket/internal/middleware"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/db_repo"
	"war_ticket/internal/repository/sqlc"
	"war_ticket/internal/usecase"
	"war_ticket/pkg"
)

func initHandler(db *sql.DB) *Inject {

	dbUr := db_repo.NewUserRepository(db)
	dbEr := db_repo.NewEventRepository(db)
	dbTr := db_repo.NewTicketRepository(db)

	sqlcQueries := sqlc.New(db)

	er := repository.NewEventRepository()
	ec := usecase.NewEventUsecase(er, dbEr)
	eh := handler.NewEventHandler(ec)

	ter := repository.NewTicketEventRepository()

	tr := repository.NewTicketRepository()
	tc := usecase.NewTicketUsecase(er, tr, ter, dbEr, dbTr, sqlcQueries)
	th := handler.NewTicketHandler(tc)

	or := repository.NewOrderRepository()
	oc := usecase.NewOrderUsecase(or, tr, db, sqlcQueries)
	oh := handler.NewOrderHandler(oc)

	// gin handler
	geh := ginhandler.NewEventHandler(ec)
	gth := ginhandler.NewTicketHandler(tc)
	goh := ginhandler.NewOrderHandler(oc)

	generateEvent(ec)
	generateTicket(tc)
	generateUser(dbUr)

	return &Inject{
		eh:  eh,
		th:  th,
		oh:  oh,
		geh: geh,
		gth: gth,
		goh: goh,
		ur:  dbUr,
	}
}

func initRouter(
	eventHandler handler.EventHandler,
	ticketHandler handler.TicketHandler,
	orderHandler handler.OrderHandler,
	userRepository repository.UserRepository,
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
	router.POST("/api/orders", middleware.AuthMiddleware(orderHandler.Create, userRepository))
	router.GET("/api/orders/length", orderHandler.GetTotalElements)

	return router
}

func initRouterGin(
	eventHandler ginhandler.EventHandler,
	ticketHandler ginhandler.TicketHandler,
	orderHandler ginhandler.OrderHandler,
	userRepository repository.UserRepository,
) *gin.Engine {
	router := gin.New()

	apiGroup := router.Group("/api")

	// event routes
	eventsGroup := apiGroup.Group("/events")
	{
		eventsGroup.GET("", eventHandler.FindAll)
		eventsGroup.POST("", eventHandler.Create)
	}

	// ticket routes
	ticketsGroup := apiGroup.Group("/tickets")
	{
		ticketsGroup.GET("", ticketHandler.FindAll)
		ticketsGroup.POST("", ticketHandler.Create)
	}

	// order routes
	ordersGroup := apiGroup.Group("/orders")
	{
		ordersGroup.GET("", orderHandler.FindAll)
		ordersGroup.POST("", middleware.AuthMiddlewareGin(userRepository), orderHandler.Create)
	}

	return router
}
