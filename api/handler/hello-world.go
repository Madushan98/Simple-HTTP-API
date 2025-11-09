package handler

import (
	"encoding/json"
	"net/http"
	"strings"
)

type MessageResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		writeError(w)
		return
	}

	firstChar := strings.ToUpper(string(name[0]))
	if firstChar >= "A" && firstChar <= "M" {
		writeSuccess(w, name)
		return
	} 

	writeError(w)		
}

func writeSuccess(w http.ResponseWriter, name string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(MessageResponse{Message: "Hello " + name})
}

func writeError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(MessageResponse{Error: "Invalid Input"})
}
