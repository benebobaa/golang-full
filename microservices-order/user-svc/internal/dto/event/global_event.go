package event

import (
	"encoding/json"
	"time"
	"user-svc/internal/dto"
	// "user-svc/internal/dto"
)

type GlobalEvent struct {
	EventID   string      `json:"event_id"`
	EventType string      `json:"event_type"`
	Timestamp time.Time   `json:"timestamp"`
	Source    string      `json:"source"`
	Action    string      `json:"action"`
	Status    string      `json:"status"`
	Payload   interface{} `json:"payload"` // Keep as interface{}
}

func (e *GlobalEvent) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func FromJSON(data []byte) (*GlobalEvent, error) {
	var event GlobalEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *GlobalEvent) GetPayloadAsUserValidateRequest() (*dto.UserValidateRequest, error) {
	payloadBytes, err := json.Marshal(e.Payload)
	if err != nil {
		return nil, err
	}
	var payload dto.UserValidateRequest
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
