package controller

import "net/http"

type HealthController struct{}

// Check godoc
// @Summary      Health Check
// @Description  Returns simple OK
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string} string "OK"
// @Router       /health [get]
func (controller *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
