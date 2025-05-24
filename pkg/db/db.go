package db

import (
    "database/sql"
    "fmt"
    "os" 
    "log"
    
    _ "github.com/go-sql-driver/mysql"
)

// Connect establishes a connection to the MySQL database
func Connect() (*sql.DB, error) {
    // Get database connection parameters from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    log.Printf("Connecting to database %s at %s:%s as user %s", dbName, dbHost, dbPort, dbUser)
    
    // Create the connection string
    dsn := fmt.Sprintf("%s:%s@unix(/var/run/mysqld/mysqld.sock)/%s?parseTime=true", 
    dbUser, dbPassword, dbName)
    
    // Open a connection to the database
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    // Test the connection
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    
    return db, nil
}
