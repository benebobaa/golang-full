package domain

type Order struct {
	Common
	Customer    string
	EventTicket []TicketEvent
}
