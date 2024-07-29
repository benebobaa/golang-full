package dto

type OrderRequest struct {
	Name    string               `json:"name" valo:"notblank"`
	Tickets []TicketOrderRequest `json:"ticket_ids" valo:"sizeMin=1,valid"`
}

type TicketOrderRequest struct {
	TicketID int `json:"ticket_id" valo:"notnil,min=1"`
	Quantity int `json:"quantity" valo:"notnil,min=1"`
}
