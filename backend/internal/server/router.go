package server

import (
	"calculator-app/backend/internal/arithmetic"
	"calculator-app/backend/internal/config"
	"calculator-app/backend/internal/expression"
	"calculator-app/backend/internal/mathops"
	"net/http"
)

// NewRouter retains the original combined server for backward compatibility.
func NewRouter(cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	arithmetic.NewHandler(arithmetic.NewService(), cfg).RegisterRoutes(mux)
	mathops.NewHandler(mathops.NewService(), cfg).RegisterRoutes(mux)
	expression.NewHandler(cfg).RegisterRoutes(mux)
	registerHealth(mux)
	return mux
}
