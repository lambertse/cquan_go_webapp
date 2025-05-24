package main

import (
	"log"

	"github.com/lambertse/cquan_go_webapp/pkg/config"
	"github.com/lambertse/cquan_go_webapp/pkg/models"
)

var appConfig *config.AppConfig 

func main() {
  appConfig, err := config.GetAppConfigFromEnv();
  if err != nil {
    log.Printf("Can not retrieve config from .env, error: %s", err)
  }

  users, err := models.GetAllUsers(appConfig.Database);
  if err != nil {
    log.Print("Can not query users")
  }
  log.Print("Users ", users)

  log.Printf("Start serving on port %s", appConfig.Port)
  appConfig.Database.Close()
}
