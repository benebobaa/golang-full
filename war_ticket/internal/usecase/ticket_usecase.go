package usecase

import (
	"war_ticket/internal/domain"
	"war_ticket/internal/domain/dto"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/repository"
)

type TicketUsecaseImpl struct {
	eventRepository       repository.EventRepository
	ticketRepository      repository.TicketRepository
	ticketEventRepository repository.TicketEventRepository
}

type TicketUsecase interface {
	Save(value *dto.TicketRequest) (*dto.TicketResponse, error)
	interfaces.Getaller[domain.Ticket]
}

func NewTicketUsecase(
	er repository.EventRepository,
	tr repository.TicketRepository,
	ter repository.TicketEventRepository,
) TicketUsecase {
	return &TicketUsecaseImpl{
		eventRepository:       er,
		ticketRepository:      tr,
		ticketEventRepository: ter,
	}
}

// Save implements TicketUsecase.
func (t *TicketUsecaseImpl) Save(value *dto.TicketRequest) (*dto.TicketResponse, error) {

	event, err := t.eventRepository.FindByID(value.EventID)

	if err != nil {
		return nil, err
	}

	ticketRequest := domain.Ticket{
		Name:  value.Name,
		Stock: value.Stock,
		Price: value.Price,
	}

	ticket, err := t.ticketRepository.Save(&ticketRequest)

	if err != nil {
		return nil, err
	}

	ticketEvent := domain.TicketEvent{
		Event:  *event,
		Ticket: *ticket,
	}

	_, err = t.ticketEventRepository.Save(&ticketEvent)

	if err != nil {
		return nil, err
	}

	return &dto.TicketResponse{
		Event:  *event,
		Ticket: *ticket,
	}, nil
}

// GetAll implements TicketUsecase.
func (t *TicketUsecaseImpl) GetAll() []domain.Ticket {
	panic("unimplemented")
}
