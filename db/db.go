package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establishes a connection to the MySQL database
func Connect() error {
  // Get database connection parameters from environment variables
  dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Configure GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Set log level to info for development
	}

	// Open a connection to the database
  var err error
  DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return err
  }

	// Get the underlying SQL database
	sqlDB, err := DB.DB()
	if err != nil {
		return err 
  }

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
  return nil
}
