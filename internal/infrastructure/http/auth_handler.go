package http

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-jwt-authentication/internal/infrastructure/http/interfaces"
	"github.com/ilyasa1211/go-jwt-authentication/internal/services"
)

type AuthHandler struct {
	Svc *services.AuthService
}

func NewAuthHandler(s *services.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	token, err := h.Svc.Login(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(&FailedResponse{
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&Response{
		Data: interfaces.LoginResponse{
			Token:     token,
			TokenType: "Bearer",
		},
	})

}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	token, err := h.Svc.Register(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(&FailedResponse{
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&Response{
		Data: interfaces.RegisterResponse{
			Token:     token,
			TokenType: "Bearer",
		},
	})
}
