package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Conn ()(*pgx.Conn, error){
	cwd, err := os.Getwd()
	if err != nil {
	    log.Fatal("Failed to get working directory:", err)
	}
	
	err = godotenv.Load(cwd + "/.env.local")
	if err != nil {
	    log.Println("Warning: .env file not found, using system environment variables.")
	}
	dbURL := os.Getenv("DATABASE_URL")
	
	db, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return db, nil
}