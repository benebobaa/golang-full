package domain

type Order struct {
	ID         int `json:"id"`
	Customer   string
	Username   string
	Tickets    []Ticket
	TotalPrice float64
	Common
}
