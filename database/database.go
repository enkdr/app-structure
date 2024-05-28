package database

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq" 
)

func InitDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }
    fmt.Println("postgres connected")
    return db, nil
}
