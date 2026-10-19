package arithmetic

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"calculator-app/backend/internal/api"
	"calculator-app/backend/internal/config"
)

func newTestMux(t *testing.T) *http.ServeMux {
	t.Helper()

	cfg, err := config.Load()
	if err != nil {
		t.Fatal(err)
	}

	service := NewService()
	handler := NewHandler(service, cfg)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return mux
}

func TestArithmeticAPI(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		body     string
		expected float64
	}{
		{"Add", api.ArithmeticAdd, `{"a":10,"b":5}`, 15},
		{"Subtract", api.ArithmeticSubtract, `{"a":10,"b":5}`, 5},
		{"Multiply", api.ArithmeticMultiply, `{"a":10,"b":5}`, 50},
		{"Divide", api.ArithmeticDivide, `{"a":10,"b":5}`, 2},
		{"ZeroOperand", api.ArithmeticAdd, `{"a":0,"b":5}`, 5},
	}

	mux := newTestMux(t)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				tc.path,
				strings.NewReader(tc.body),
			)

			recorder := httptest.NewRecorder()
			mux.ServeHTTP(recorder, req)

			if recorder.Code != http.StatusOK {
				t.Fatalf(
					"expected 200, got %d: %s",
					recorder.Code,
					recorder.Body.String(),
				)
			}

			var response CalculateResponse

			if err := json.Unmarshal(
				recorder.Body.Bytes(),
				&response,
			); err != nil {
				t.Fatal(err)
			}

			if response.Result != tc.expected {
				t.Fatalf(
					"expected %v, got %v",
					tc.expected,
					response.Result,
				)
			}
		})
	}
}

func TestArithmeticAPIErrors(t *testing.T) {
	tests := []struct {
		name string
		path string
		body string
	}{
		{
			"DivisionByZero",
			api.ArithmeticDivide,
			`{"a":10,"b":0}`,
		},
		{
			"MissingOperand",
			api.ArithmeticAdd,
			`{"a":10}`,
		},
		{
			"MalformedJSON",
			api.ArithmeticAdd,
			`{"a":10,"b":}`,
		},
		{
			"UnknownField",
			api.ArithmeticAdd,
			`{"a":10,"b":5,"extra":1}`,
		},
	}

	mux := newTestMux(t)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				tc.path,
				strings.NewReader(tc.body),
			)

			recorder := httptest.NewRecorder()
			mux.ServeHTTP(recorder, req)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf(
					"expected 400, got %d: %s",
					recorder.Code,
					recorder.Body.String(),
				)
			}
		})
	}
}

func TestArithmeticAPIRequestLimit(t *testing.T) {
	cfg := config.Config{
		MaxRequestBytes: 10,
	}

	handler := NewHandler(NewService(), cfg)
	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(
		http.MethodPost,
		api.ArithmeticAdd,
		strings.NewReader(`{"a":100,"b":200}`),
	)

	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected 400, got %d",
			recorder.Code,
		)
	}
}

func TestArithmeticAPIMethodNotAllowed(t *testing.T) {
	mux := newTestMux(t)

	req := httptest.NewRequest(
		http.MethodGet,
		api.ArithmeticAdd,
		nil,
	)

	recorder := httptest.NewRecorder()
	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected 405, got %d",
			recorder.Code,
		)
	}
}
