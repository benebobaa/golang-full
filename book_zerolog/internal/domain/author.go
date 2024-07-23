package domain

type Author struct {
	Name string `json:"name" valo:"notblank,sizeMin=8"`
	City string `json:"city" valo:"notblank"`
}
