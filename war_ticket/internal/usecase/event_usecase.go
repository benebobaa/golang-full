package usecase

import (
	"war_ticket/internal/domain"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/repository"
)

type EventUsecaseImpl struct {
	eventRepository repository.EventRepository
}

type EventUsecase interface {
	interfaces.Getaller[domain.Event]
	interfaces.Saver[domain.Event]
}

func NewEventUsecase(eventRepository repository.EventRepository) EventUsecase {
	return &EventUsecaseImpl{
		eventRepository: eventRepository,
	}
}

// GetAll implements EventUsecase.
func (e *EventUsecaseImpl) GetAll() []domain.Event {
	return e.eventRepository.GetAll()
}

// Save implements EventUsecase.
func (e *EventUsecaseImpl) Save(value *domain.Event) (*domain.Event, error) {

	event, err := e.eventRepository.Save(value)

	if err != nil {
		return nil, err
	}

	return event, nil
}
