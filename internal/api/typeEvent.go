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
	Status    string `json:"status"`
	EventID   string `json:"id,omitempty"`
	Timestamp string `json:"timestamp"`
}

func (e *EventRequest) Validate() error {
	if e.EventType == "" {
		return fmt.Errorf("Event type cannot be empty")
	}

	if len(e.Payload) == 0 {
		return fmt.Errorf("Payload cannot be empty")
	}
	return nil
}
