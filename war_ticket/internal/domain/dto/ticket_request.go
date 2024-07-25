package dto

import "war_ticket/internal/domain"

type TicketRequest struct {
	EventID int
	domain.Ticket
}
