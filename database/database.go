package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func ConnectDB() *pgx.Conn {
	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	database, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil || !ok {
		fmt.Printf("[Database] Unable to connect: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[Database] Successfully connected.")
	return database
}

func DisconnectDB(database *pgx.Conn) {
	err := database.Close(context.Background())
	if err != nil {
		fmt.Printf("[Database] Unable to close: %v\n", err)
	}
}
