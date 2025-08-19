package controller

import (
	"encoding/json"
	"net/http"

	"github.com/berpergian/chi_learning/account_service/message"
	"github.com/berpergian/chi_learning/account_service/service"
	"github.com/berpergian/chi_learning/shared/config"
)

type AccountController struct {
	AccountService *service.AccountService
	Env            *config.Env
}

func RegisterAccountController(env *config.Env, accountService *service.AccountService) *AccountController {
	return &AccountController{AccountService: accountService, Env: env}
}

// RegisterOrLogin godoc
// @Summary      Register or Login
// @Description  Register a new account if not exist, otherwise login
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body message.LoginRequest true "Login Request"
// @Success      200  {object} message.LoginResponse
// @Failure      400  {string} string "Bad Request"
// @Router       /registerOrLogin [post]
func (controller *AccountController) RegisterOrLogin(w http.ResponseWriter, r *http.Request) {
	var req message.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp, err := controller.AccountService.RegisterOrLogin(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
