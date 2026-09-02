package api

import (
	"encoding/json"
	"fmt"
)

type EventRequest struct {
	EventType string `json:"eventType"`
	//el segudno string le indica al lenguaje que tipo eseprar en esa clave es QUE espera
	Payload json.RawMessage `json:"payload"`
}

type EventResponse struct {
	Status  string          `json:"status"`
	Payload json.RawMessage `json:"payload"`
}

func (e *EventRequest) ValidateMethodPost() error {
	if e.EventType != "POST" {
		return fmt.Errorf("Invalid event type: %s", e.EventType)
	}
	return nil
}

func (e *EventRequest) ValidatePayload() error {
	if len(e.Payload) == 0 {
		return fmt.Errorf("Payload cannot be empty")
	}
	return nil
}
