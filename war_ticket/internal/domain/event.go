package domain

type Event struct {
	ID       int    `json:"id"`
	Name     string `json:"name" valo:"notblank"`
	Location string `json:"location" valo:"notblank"`
	Common
}
