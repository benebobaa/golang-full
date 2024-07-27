package dto

import "war_ticket/internal/domain"

type TicketRequest struct {
	EventID       int `json:"event_id" valo:"notnil"`
	domain.Ticket `valo:"valid"`
}
