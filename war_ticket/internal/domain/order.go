package domain

type Order struct {
	ID         int      `json:"id"`
	Customer   string   `json:"customer"`
	Username   string   `json:"username"`
	Tickets    []Ticket `json:"tickets"`
	TotalPrice float64  `json:"total_price"`
	Common
}
