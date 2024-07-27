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
	interfaces.Getter[domain.Ticket]
	GetAllWithEvent() ([]dto.TicketEventResponse, error)
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
		EventID:  event.ID,
		TicketID: ticket.ID,
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
	return t.ticketRepository.GetAll()
}

// GetAllWithEvent implements TicketUsecase.
func (t *TicketUsecaseImpl) GetAllWithEvent() ([]dto.TicketEventResponse, error) {
	tEvents := t.ticketEventRepository.GetAll()

	eventMap := make(map[int]*dto.TicketEventResponse)

	for _, v := range tEvents {
		if _, exists := eventMap[v.EventID]; !exists {
			event, err := t.eventRepository.FindByID(v.EventID)
			if err != nil {
				return nil, err
			}
			eventMap[v.EventID] = &dto.TicketEventResponse{
				Event: *event,
			}
		}

		ticket, err := t.ticketRepository.FindByID(v.TicketID)
		if err != nil {
			return nil, err
		}

		eventMap[v.EventID].Tickets = append(eventMap[v.EventID].Tickets, *ticket)
	}

	result := make([]dto.TicketEventResponse, 0, len(eventMap))
	for _, resp := range eventMap {
		result = append(result, *resp)
	}

	return result, nil
}
