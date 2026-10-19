package server

import (
	"calculator-app/backend/internal/config"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndependentServiceRoutes(t *testing.T) {
	cfg := config.Config{MaxRequestBytes: 4096, MaxExpressionLength: 1024}
	tests := []struct {
		name     string
		router   http.Handler
		allowed  string
		body     string
		excluded string
	}{
		{"arithmetic", NewArithmeticRouter(cfg), "/api/v1/arithmetic/add", `{"a":2,"b":3}`, "/api/v1/math/power"},
		{"math", NewMathRouter(cfg), "/api/v1/math/power", `{"a":2,"b":3}`, "/api/v1/expressions/evaluate"},
		{"expression", NewExpressionRouter(cfg), "/api/v1/expressions/evaluate", `{"expression":"2+3"}`, "/api/v1/arithmetic/add"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			health := httptest.NewRecorder()
			tt.router.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/health", nil))
			if health.Code != 200 {
				t.Fatalf("health status %d", health.Code)
			}
			req := httptest.NewRequest(http.MethodPost, tt.allowed, strings.NewReader(tt.body))
			res := httptest.NewRecorder()
			tt.router.ServeHTTP(res, req)
			if res.Code != 200 {
				t.Fatalf("operation status %d: %s", res.Code, res.Body.String())
			}
			absent := httptest.NewRecorder()
			tt.router.ServeHTTP(absent, httptest.NewRequest(http.MethodPost, tt.excluded, strings.NewReader(tt.body)))
			if absent.Code != 404 {
				t.Fatalf("cross-service endpoint returned %d", absent.Code)
			}
		})
	}
}
