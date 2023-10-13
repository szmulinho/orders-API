package main

import (
	"fmt"
	"github.com/szmulinho/common/utils"
	"github.com/szmulinho/orders/database"
	"github.com/szmulinho/orders/internal/server"
	"log"
)

func main() {
	fmt.Println("Staring the application...")
	defer fmt.Println("Closing the application...")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}

	ctx, _, wait := utils.Gracefully()

	server.Run(ctx, db)

	wait()
}
