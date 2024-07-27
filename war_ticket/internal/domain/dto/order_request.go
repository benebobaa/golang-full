package dto

type OrderRequest struct {
	Name    string               `json:"name" valo:"notblank"`
	Tickets []TicketOrderRequest `json:"ticket_ids" valo:"notnil, valid"`
}

type TicketOrderRequest struct {
	TicketID int `json:"ticket_id" valo:"notnil"`
	Quantity int `json:"quantity" valo:"notnil"`
}
