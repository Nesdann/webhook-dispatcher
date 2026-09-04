package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        string
	EventType string
	Body      json.RawMessage
	CreatedAt time.Time
}

func NewEvent(eventype string, payload json.RawMessage) Event {
	var res Event
	res.ID = uuid.New().String() //este luego tamiben es cual se debe delvovler en la respuesta

	res.EventType = eventype
	res.Body = payload
	res.CreatedAt = time.Now().UTC()
	return res

}
