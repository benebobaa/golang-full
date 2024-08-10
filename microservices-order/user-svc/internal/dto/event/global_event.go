package event

import (
	"encoding/json"
	"time"
)

type GlobalEvent[T any] struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
	Action    string    `json:"action"`
	Status    string    `json:"status"`
	Payload   T         `json:"payload"`
}

func (ge GlobalEvent[T]) ToJSON() ([]byte, error) {
	return json.Marshal(ge)
}

func FromJSON[T any](data []byte) (GlobalEvent[T], error) {
	var ge GlobalEvent[T]
	err := json.Unmarshal(data, &ge)
	return ge, err
}

func NewEvent[T any]() GlobalEvent[T] {
	return GlobalEvent[T]{
		EventID:   "",
		EventType: "",
		Timestamp: time.Time{},
		Source:    "user-svc",
		Action:    "validate",
		Status:    "",
		Payload:   *new(T),
	}
}
