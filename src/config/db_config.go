package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL") // This will be your NeonDB connection string
	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	// Optional: Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL (NeonDB)")
}
