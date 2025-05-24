package main

import (
	"log"

	"github.com/lambertse/cquan_go_webapp/db"
	"github.com/lambertse/cquan_go_webapp/pkg/config"
	"github.com/lambertse/cquan_go_webapp/services"
)

var appConfig *config.AppConfig 

func main() {
  appConfig, err := config.GetAppConfigFromEnv();
  if err != nil {
    log.Printf("Can not retrieve config from .env, error: %s", err)
  }

  err = db.Connect();
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }

  // Migrate and seed data
  // err = db.MigrateAndSeed()
  // if err != nil {
  //   log.Fatalf("Migrate and seed failed: %v", err)
  // }
  //

  userService := services.NewUserService()
  user, err := userService.GetAllUsers()
  if err != nil {
    log.Fatalf("Failed to get users: %v", err)
  }
  // Print users detail
  for _, u := range user {
    log.Printf("User: ID=%d, Username=%s", u.ID, u.Username)
  }
  log.Printf("Start serving on port %s", appConfig.Port)
}
