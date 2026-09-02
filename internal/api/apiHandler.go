package api

import (
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := req.ValidateMethodPost(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.ValidatePayload(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respuesta := EventResponse{
		Status:  "success",
		Payload: json.RawMessage(`{"message":"bien curl mandado bien"}`),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(respuesta); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
