package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Caso feliz: request bien formado -> 202 Accepted + EventID generado.
func TestHandler_ValidEvent_Returns202(t *testing.T) {
	body := `{"eventType":"order.created","payload":{"orderId":123,"total":99.5}}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("esperaba status 202, recibí %d. Body: %s", rec.Code, rec.Body.String())
	}

	var resp EventResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("la respuesta no es JSON válido: %v", err)
	}

	if resp.Status != "accepted" {
		t.Errorf("esperaba status=\"accepted\", recibí %q", resp.Status)
	}
	if resp.EventID == "" {
		t.Error("esperaba un EventID generado (UUID), vino vacío")
	}
	if resp.Timestamp == "" {
		t.Error("esperaba un Timestamp, vino vacío")
	}

	// Content-Type de la respuesta también es parte del contrato de la API.
	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("esperaba Content-Type application/json, recibí %q", ct)
	}
}

// Dos requests válidos consecutivos deben generar EventIDs distintos.
// Esto prueba que uuid.New() realmente se está llamando por request,
// y no hay un ID hardcodeado o cacheado por error.
func TestHandler_ValidEvent_GeneratesUniqueEventIDs(t *testing.T) {
	body := `{"eventType":"order.created","payload":{"orderId":1}}`

	req1 := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	rec1 := httptest.NewRecorder()
	Handler(rec1, req1)

	req2 := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	rec2 := httptest.NewRecorder()
	Handler(rec2, req2)

	var resp1, resp2 EventResponse
	json.Unmarshal(rec1.Body.Bytes(), &resp1)
	json.Unmarshal(rec2.Body.Bytes(), &resp2)

	if resp1.EventID == resp2.EventID {
		t.Errorf("dos requests distintos generaron el mismo EventID: %s", resp1.EventID)
	}
}

// Método distinto a POST -> 405 Method Not Allowed.
func TestHandler_WrongMethod_Returns405(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/event", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("esperaba status 405, recibí %d", rec.Code)
	}
}

// JSON malformado (body cortado/inválido) -> 400 Bad Request.
func TestHandler_MalformedJSON_Returns400(t *testing.T) {
	body := `{"eventType": "order.created", "payload":`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba status 400, recibí %d", rec.Code)
	}
}

// eventType vacío -> 400 Bad Request.
// Este es el test que confirma que el bug de ValidateMethodPost quedó resuelto:
// la validación ahora es sobre el campo de negocio, no sobre el método HTTP.
func TestHandler_EmptyEventType_Returns400(t *testing.T) {
	body := `{"eventType":"","payload":{"x":1}}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba status 400, recibí %d. Body: %s", rec.Code, rec.Body.String())
	}
}

// eventType ausente directamente en el JSON (no solo vacío) -> también 400.
// json.Decode deja el campo en su zero-value (""), así que debería
// comportarse igual que el caso anterior. Vale la pena probarlo explícito.
func TestHandler_MissingEventType_Returns400(t *testing.T) {
	body := `{"payload":{"x":1}}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba status 400, recibí %d", rec.Code)
	}
}

// payload vacío/ausente -> 400 Bad Request.
func TestHandler_EmptyPayload_Returns400(t *testing.T) {
	body := `{"eventType":"order.created"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/event", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	Handler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("esperaba status 400, recibí %d", rec.Code)
	}
}
