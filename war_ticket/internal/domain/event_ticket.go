package domain

type TicketEvent struct {
	Common
	Event  Event
	Ticket Ticket
}
