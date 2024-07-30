package main

import (
	"context"
	"database/sql"
	"log"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/handler"
	"war_ticket/internal/middleware"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/db_repo"
	"war_ticket/internal/repository/sqlc"
	"war_ticket/internal/usecase"
	"war_ticket/pkg"

	"github.com/google/uuid"
)

func initHandler(db *sql.DB) (handler.EventHandler, handler.TicketHandler, handler.OrderHandler, repository.UserRepository) {

	dbUr := db_repo.NewUserRepository(db)
	dbEr := db_repo.NewEventRepository(db)
	dbTr := db_repo.NewTicketRepository(db)

	sqlc := sqlc.New(db)

	er := repository.NewEventRepository()
	ec := usecase.NewEventUsecase(er, dbEr)
	eh := handler.NewEventHandler(ec)

	ter := repository.NewTicketEventRepository()

	tr := repository.NewTicketRepository()
	tc := usecase.NewTicketUsecase(er, tr, ter, dbEr, dbTr, sqlc)
	th := handler.NewTicketHandler(tc)

	or := repository.NewOrderRepository()
	oc := usecase.NewOrderUsecase(or, tr)
	oh := handler.NewOrderHandler(oc)

	ur := repository.NewUserRepository()

	generateEvent(ec)
	generateTicket(tc)
	generateUser(ur)
	generateUser(dbUr)

	return eh, th, oh, ur
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

func generateEvent(ec usecase.EventUsecase) {
	event1 := domain.Event{
		Name:     "Lomba joget",
		Location: "Jaksel",
	}

	event2 := domain.Event{
		Name:     "Konser Nyanyi",
		Location: "Blok M",
	}

	ec.Save(context.Background(), &event1)
	ec.Save(context.Background(), &event2)
}

func generateTicket(tc usecase.TicketUsecase) {
	ticket1 := dto.TicketRequest{
		EventID: 1,
		Ticket: domain.Ticket{
			Name:  "VIP 1",
			Stock: 10,
			Price: 5000,
		},
	}
	ticket2 := dto.TicketRequest{
		EventID: 1,
		Ticket: domain.Ticket{
			Name:  "CAT 1",
			Stock: 100,
			Price: 250,
		},
	}
	tc.Save(context.Background(), &ticket1)
	tc.Save(context.Background(), &ticket2)

	ticket1.EventID = 2
	ticket2.EventID = 2

	tc.Save(context.Background(), &ticket1)
	tc.Save(context.Background(), &ticket2)
}

func generateUser(userRepository repository.UserRepository) {

	user1 := domain.User{
		ApiKey:   uuid.NewString(),
		Username: "kapallaut",
	}
	user2 := domain.User{
		ApiKey:   uuid.NewString(),
		Username: "beneboba",
	}

	userRepository.Save(context.Background(), &user1)
	userRepository.Save(context.Background(), &user2)
	log.Println("user 1 :: ", user1)
	log.Println("user 2 :: ", user2)
}
