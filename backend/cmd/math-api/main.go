package main

import (
	"calculator-app/backend/internal/server"
	"log"
)

func main() {
	if err := server.Run("math-api", 8082, server.NewMathRouter); err != nil {
		log.Fatal(err)
	}
}
