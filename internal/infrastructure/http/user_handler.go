package http

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasa1211/go-jwt-authentication/internal/services"
)

type UserHandler struct {
	Svc *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.FindAll()

	resp := &Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.FindById(r)

	resp := &Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.Create(r)

	resp := &Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.UpdateById(r)

	resp := &Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	res := h.Svc.DeleteById(r)

	resp := &Response{
		Data: res,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
