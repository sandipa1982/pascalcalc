package main

import (
	"calculator-app/backend/internal/server"
	"log"
)

func main() {
	if err := server.Run("arithmetic-api", 8081, server.NewArithmeticRouter); err != nil {
		log.Fatal(err)
	}
}
