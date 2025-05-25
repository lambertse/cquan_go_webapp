package rest_handlers

import (
	"net/http"

	"github.com/lambertse/cquan_go_webapp/internal/services"
)

type UserHandler struct {
  UserService *services.UserService
}

func NewUserHandler(service *services.UserService) (*UserHandler){
  handler := UserHandler{
    UserService: service,
  }
  return &handler
} 

func (m *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
  users, err := m.UserService.GetAllUsers()
  if err != nil {
    http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
  }
  for _, user := range users {
    _, err := w.Write([]byte(user.Username + "\n"))
    if err != nil {
      http.Error(w, "Failed to write user", http.StatusInternalServerError)
      return
    }
  }
}


