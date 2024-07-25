package domain

type Event struct {
	Name     string `json:"name" valo:"notblank"`
	Location string `json:"location" valo:"notblank"`
	Common
}
