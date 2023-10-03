package main

import (
	"github.com/szmulinho/orders/cmd/server"
	"github.com/szmulinho/orders/internal/database"
)

func main() {
	database.Connect()

	server.Run()
}
