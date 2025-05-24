package models

import (
    "database/sql"
    "time"
    // "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID        int64     `json:"id"`
    Username  string    `json:"username"`
    Password  string    `json:"-"` // Use "-" to exclude from JSON responses
    CreatedAt time.Time `json:"created_at"`
}

func GetAllUsers(db *sql.DB) ([]*User, error) {
    query := "SELECT id, username, created_at FROM users"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    users := make([]*User, 0)
    for rows.Next() {
        user := &User{}
        if err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    if err := rows.Err(); err != nil {
        return nil, err
    }
    
    return users, nil
}
