package domain

type Order struct {
	Customer   string
	Tickets    []Ticket
	TotalPrice float64
	Common
}
