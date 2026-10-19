package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"calculator-app/backend/internal/api"
	"calculator-app/backend/internal/config"
)

func TestRouterHealth(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	router := NewRouter(cfg)

	req := httptest.NewRequest(
		http.MethodGet,
		api.Health,
		nil,
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected 200, got %d",
			recorder.Code,
		)
	}
}

func TestRouterArithmetic(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	router := NewRouter(cfg)

	req := httptest.NewRequest(
		http.MethodPost,
		api.ArithmeticAdd,
		strings.NewReader(`{"a":10,"b":5}`),
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected 200, got %d: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}
}

func TestRouterMath(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	router := NewRouter(cfg)

	req := httptest.NewRequest(
		http.MethodPost,
		api.MathSquareRoot,
		strings.NewReader(`{"a":81}`),
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected 200, got %d: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}
}

func TestRouterExpression(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	router := NewRouter(cfg)

	req := httptest.NewRequest(
		http.MethodPost,
		api.ExpressionEvaluate,
		strings.NewReader(
			`{"expression":"20 + 5 * 3 - 1"}`,
		),
	)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected 200, got %d: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}
}
