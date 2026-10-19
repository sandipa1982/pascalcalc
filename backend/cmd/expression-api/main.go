package main

import (
	"calculator-app/backend/internal/server"
	"log"
)

func main() {
	if err := server.Run("expression-api", 8083, server.NewExpressionRouter); err != nil {
		log.Fatal(err)
	}
}
