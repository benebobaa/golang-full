package domain

type Ticket struct {
	Name  string  `json:"name" valo:"notblank"`
	Stock int     `json:"stock" valo:"min=1"`
	Price float64 `json:"price" valo:"min=1"`
	Common
}
