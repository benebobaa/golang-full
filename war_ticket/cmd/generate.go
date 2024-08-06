package main

import (
	"context"
	"github.com/google/uuid"
	"log"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/repository"
	"war_ticket/internal/usecase"
)

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
