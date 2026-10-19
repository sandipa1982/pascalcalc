package mathops

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

	handler := NewHandler(NewService(), cfg)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	return mux
}

func TestMathAPI(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		body     string
		expected float64
	}{
		{
			"Power",
			api.MathPower,
			`{"a":2,"b":3}`,
			8,
		},
		{
			"SquareRoot",
			api.MathSquareRoot,
			`{"a":81}`,
			9,
		},
		{
			"Percentage",
			api.MathPercentage,
			`{"a":20,"b":150}`,
			30,
		},
		{
			"ZeroOperand",
			api.MathPower,
			`{"a":0,"b":3}`,
			0,
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

			if recorder.Code != http.StatusOK {
				t.Fatalf(
					"expected 200, got %d: %s",
					recorder.Code,
					recorder.Body.String(),
				)
			}

			var response MathResponse

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

func TestMathAPIErrors(t *testing.T) {
	tests := []struct {
		name string
		path string
		body string
	}{
		{
			"NegativeSquareRoot",
			api.MathSquareRoot,
			`{"a":-9}`,
		},
		{
			"MissingOperand",
			api.MathPower,
			`{"a":2}`,
		},
		{
			"MalformedJSON",
			api.MathPower,
			`{"a":2,"b":}`,
		},
		{
			"UnknownField",
			api.MathPower,
			`{"a":2,"b":3,"extra":1}`,
		},
		{
			"PowerOverflow",
			api.MathPower,
			`{"a":1e308,"b":2}`,
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

func TestMathAPIRequestLimit(t *testing.T) {
	cfg := config.Config{
		MaxRequestBytes: 10,
	}

	handler := NewHandler(NewService(), cfg)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	req := httptest.NewRequest(
		http.MethodPost,
		api.MathPower,
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

func TestMathAPIMethodNotAllowed(t *testing.T) {
	mux := newTestMux(t)

	req := httptest.NewRequest(
		http.MethodGet,
		api.MathSquareRoot,
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
