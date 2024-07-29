package domain

type Order struct {
	Customer   string
	Username   string
	Tickets    []Ticket
	TotalPrice float64
	Common
}
