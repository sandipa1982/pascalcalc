package server

import (
	"calculator-app/backend/internal/api"
	"calculator-app/backend/internal/arithmetic"
	"calculator-app/backend/internal/config"
	"calculator-app/backend/internal/expression"
	"calculator-app/backend/internal/mathops"
	"net/http"
)

// NewArithmeticRouter exposes only arithmetic operations and health.
func NewArithmeticRouter(cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	arithmetic.NewHandler(arithmetic.NewService(), cfg).RegisterRoutes(mux)
	registerHealth(mux)
	return mux
}

// NewMathRouter exposes only advanced mathematical operations and health.
func NewMathRouter(cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	mathops.NewHandler(mathops.NewService(), cfg).RegisterRoutes(mux)
	registerHealth(mux)
	return mux
}

// NewExpressionRouter exposes only expression evaluation and health.
func NewExpressionRouter(cfg config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	expression.NewHandler(cfg).RegisterRoutes(mux)
	registerHealth(mux)
	return mux
}

func registerHealth(mux *http.ServeMux) {
	mux.HandleFunc(api.Health, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})
}
