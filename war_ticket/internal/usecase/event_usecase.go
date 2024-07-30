package usecase

import (
	"context"
	"war_ticket/internal/domain"
	"war_ticket/internal/interfaces"
	"war_ticket/internal/repository"
	"war_ticket/internal/repository/db_repo"
)

type EventUsecaseImpl struct {
	eventRepository repository.EventRepository
	eventdb         db_repo.EventRepository
}

type EventUsecase interface {
	interfaces.Getter[domain.Event]
	interfaces.Saver[domain.Event]
}

func NewEventUsecase(
	eventRepository repository.EventRepository,
	eventdb db_repo.EventRepository,
) EventUsecase {
	return &EventUsecaseImpl{
		eventRepository: eventRepository,
		eventdb:         eventdb,
	}
}

// GetAll implements EventUsecase.
func (e *EventUsecaseImpl) GetAll() []domain.Event {
	// return e.eventRepository.GetAll()
	return e.eventdb.GetAll()
}

// Save implements EventUsecase.
func (e *EventUsecaseImpl) Save(ctx context.Context, value *domain.Event) (*domain.Event, error) {
	// event, err := e.eventRepository.Save(ctx, value)

	event, err := e.eventdb.Save(ctx, value)

	if err != nil {
		return nil, err
	}

	return event, nil
}
