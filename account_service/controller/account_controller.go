package controller

import (
	"encoding/json"
	"net/http"

	"github.com/berpergian/chi_learning/account_service/message"
	"github.com/berpergian/chi_learning/account_service/service"
	"github.com/berpergian/chi_learning/shared/config"
	"github.com/berpergian/chi_learning/shared/helper"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type AccountController struct {
	AccountService *service.AccountService
	Env            *config.Env
	Validate       *validator.Validate
}

func RegisterAccountController(env *config.Env, accountService *service.AccountService, validate *validator.Validate) *AccountController {
	return &AccountController{AccountService: accountService, Env: env, Validate: validate}
}

// Register godoc
// @Summary      Register
// @Description  Register a new account
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body message.RegisterRequest true "Register Request"
// @Success      200  {object} message.RegisterResponse
// @Failure      400  {object} helper.ProblemDetails
// @Failure      409  {object} helper.ProblemDetails
// @Router       /register [post]
func (controller *AccountController) Register(w http.ResponseWriter, r *http.Request) {
	var req message.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteProblem(w, r, helper.ProblemDetails{
			Title:     "Invalid JSON",
			Status:    http.StatusBadRequest,
			Detail:    "Request body could not be parsed",
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		validationErrors := make(map[string]map[string]string)
		for _, fe := range err.(validator.ValidationErrors) {
			field := fe.Field()

			if _, exists := validationErrors[field]; !exists {
				validationErrors[field] = map[string]string{
					"rule":  fe.Tag(),
					"value": fe.Param(),
				}
			}
		}

		helper.WriteProblem(w, r, helper.ProblemDetails{
			Title:     "Validation failed",
			Status:    http.StatusBadRequest,
			Detail:    "One or more validation errors occurred.",
			Errors:    validationErrors,
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	resp, err := controller.AccountService.Register(r.Context(), req)
	if err != nil {
		helper.WriteProblem(w, r, helper.ProblemDetails{
			Title:     "Register account failed",
			Status:    http.StatusConflict,
			Detail:    err.Error(),
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Login godoc
// @Summary      Login
// @Description  Login existing account
// @Tags         account
// @Accept       json
// @Produce      json
// @Param        request body message.LoginRequest true "Login Request"
// @Success      200  {object} message.LoginResponse
// @Failure      400  {object} helper.ProblemDetails
// @Failure      401  {object} helper.ProblemDetails
// @Router       /login [post]
func (controller *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	var req message.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.WriteProblem(w, r, helper.ProblemDetails{
			Title:     "Invalid JSON",
			Status:    http.StatusBadRequest,
			Detail:    "Request body could not be parsed",
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	if err := controller.Validate.Struct(req); err != nil {
		validationErrors := make(map[string]map[string]string)
		for _, fe := range err.(validator.ValidationErrors) {
			field := fe.Field()

			if _, exists := validationErrors[field]; !exists {
				validationErrors[field] = map[string]string{
					"rule":  fe.Tag(),
					"value": fe.Param(),
				}
			}
		}

		helper.WriteProblem(w, r, helper.ProblemDetails{
			Title:     "Validation failed",
			Status:    http.StatusBadRequest,
			Detail:    "One or more validation errors occurred.",
			Errors:    validationErrors,
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	resp, err := controller.AccountService.Login(r.Context(), req)
	if err != nil {
		helper.WriteProblem(w, r, helper.ProblemDetails{
			Title:     "Login account failed",
			Status:    http.StatusUnauthorized,
			Detail:    err.Error(),
			RequestID: r.Context().Value(middleware.RequestIDKey),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
