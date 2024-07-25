package dto

import "war_ticket/internal/domain"

type TicketResponse struct {
	Event  domain.Event
	Ticket domain.Ticket
}
