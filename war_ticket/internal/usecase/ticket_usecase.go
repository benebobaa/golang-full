package usecase

import (
	"context"
	"log"
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/db_repo"
	"war_ticket/internal/repository/sqlc"
)

type TicketUsecaseImpl struct {
	eventRepository       repository.EventRepository
	ticketRepository      repository.TicketRepository
	ticketEventRepository repository.TicketEventRepository
	eventDbRepo           db_repo.EventRepository
	ticketDbRepo          db_repo.TicketRepository
	sqlcQueries           sqlc.Querier
}

type TicketUsecase interface {
	Save(ctx context.Context, value *dto.TicketRequest) (*dto.TicketResponse, error)
	interfaces.Getter[domain.Ticket]
	GetAllWithEvent() ([]dto.TicketEventResponse, error)
}

func NewTicketUsecase(
	er repository.EventRepository,
	tr repository.TicketRepository,
	ter repository.TicketEventRepository,
	edr db_repo.EventRepository,
	tdr db_repo.TicketRepository,
	sqlcQueries sqlc.Querier,
) TicketUsecase {
	return &TicketUsecaseImpl{
		eventRepository:       er,
		ticketRepository:      tr,
		ticketEventRepository: ter,
		eventDbRepo:           edr,
		ticketDbRepo:          tdr,
		sqlcQueries:           sqlcQueries,
	}
}

// Save implements TicketUsecase.
func (t *TicketUsecaseImpl) Save(ctx context.Context, value *dto.TicketRequest) (*dto.TicketResponse, error) {

	log.Println("request :: ", value)

	event, err := t.eventDbRepo.FindByID(value.EventID)

	if err != nil {
		return nil, err
	}

	ticket, err := t.sqlcQueries.CreateTicket(
		ctx,
		sqlc.CreateTicketParams{
			Name:  value.Name,
			Stock: int32(value.Stock),
			Price: value.Price,
		},
	)

	if err != nil {
		return nil, err
	}

	err = t.sqlcQueries.CreateTicketEvent(
		ctx,
		sqlc.CreateTicketEventParams{
			EventID:  int32(event.ID),
			TicketID: int32(ticket.ID),
		},
	)

	if err != nil {
		return nil, err
	}

	return &dto.TicketResponse{
		Event:  *event,
		Ticket: ticket,
	}, nil
}

// GetAll implements TicketUsecase.
func (t *TicketUsecaseImpl) GetAll() []domain.Ticket {
	// result, _ := t.sqlcQueries.ListTickets(context.Background())
	// return result
	result, _ := t.sqlcQueries.ListTickets(context.Background())
	return result
}

// GetAllWithEvent implements TicketUsecase.
func (t *TicketUsecaseImpl) GetAllWithEvent() ([]dto.TicketEventResponse, error) {
	result, err := t.ticketDbRepo.ListTicketsWithEvents()

	if err != nil {
		return nil, err
	}

	return result, nil
}
