package dto

import "war_ticket/internal/domain"

type TicketEventResponse struct {
	Event   domain.Event    `json:"event"`
	Tickets []domain.Ticket `json:"tickets"`
}
