package helper

import (
	"encoding/json"
	"net/http"
)

type ProblemDetails struct {
	Type      string                       `json:"type,omitempty"`
	Title     string                       `json:"title"`
	Status    int                          `json:"status"`
	Detail    string                       `json:"detail,omitempty"`
	Instance  string                       `json:"instance,omitempty"`
	Errors    map[string]map[string]string `json:"errors,omitempty"`
	RequestID any                          `json:"requestId,omitempty"`
}

func WriteProblem(w http.ResponseWriter, r *http.Request, pd ProblemDetails) {
	if pd.Instance == "" {
		pd.Instance = r.Method + " " + r.URL.Path
	}

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(pd.Status)
	_ = json.NewEncoder(w).Encode(pd)
}
