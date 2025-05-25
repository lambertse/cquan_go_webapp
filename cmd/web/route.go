package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lambertse/cquan_go_webapp/internal/services"
	rest_handlers "github.com/lambertse/cquan_go_webapp/internal/transport/rest/handlers"
)

func route() http.Handler {
  mux := chi.NewRouter()

  // Create services
  userService := services.NewUserService()
  //

  // Create Handler
  userHandler := rest_handlers.NewUserHandler(userService) 
  //

  //  Routing
  mux.Get("/users", userHandler.GetAllUsers)
  //

  return mux
}
