package config

import (
	"database/sql"
	"os"
  "log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/lambertse/cquan_go_webapp/pkg/db"
)

type AppConfig struct {
  Port         string `env:"PORT" envDefault:": "8080"`
  DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://user:password@localhost:5432/mydb"`
  LogLevel    string `env:"LOG_LEVEL" envDefault:"info`
  Database    *sql.DB `env:"-"` 
}

func GetAppConfigFromEnv() (*AppConfig, error) {
  var config AppConfig
  // Get the directory of the current file
	_, filename, _, _ := runtime.Caller(0) 
	dir := filepath.Dir(filename)
	
	// Load .env file from project root
	err := godotenv.Load(filepath.Join(dir, "..", "..", ".env"))
  if err != nil {
    return &config, err
  }
  config.Port = os.Getenv("PORT")

  database, err := db.Connect()
  if err != nil {
      log.Fatalf("Failed to connect to database: %v", err)
  }  
  config.Database = database 
  return &config, nil
}
