package main

import (
	"context"
	"fmt"
	"log"

	"github.com/adudumayo/go-pool-app/internal/db"
)

func main() {
	conn, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to PostgreSQL successfully!")
}
