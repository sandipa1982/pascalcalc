package expression

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

	mux := http.NewServeMux()

	NewHandler(cfg).RegisterRoutes(mux)

	return mux
}

func TestExpressionAPI(t *testing.T) {
	mux := newTestMux(t)

	req := httptest.NewRequest(
		http.MethodPost,
		api.ExpressionEvaluate,
		strings.NewReader(
			`{"expression":"20 + 5 * 3 - 1"}`,
		),
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

	var response EvaluateResponse

	if err := json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	); err != nil {
		t.Fatal(err)
	}

	if response.Result != 34 {
		t.Fatalf(
			"expected 34, got %v",
			response.Result,
		)
	}
}

func TestExpressionAPIErrors(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{"EmptyExpression", `{"expression":""}`},
		{"DivisionByZero", `{"expression":"10 / 0"}`},
		{"InvalidSyntax", `{"expression":"10 +"}`},
		{"InvalidCharacter", `{"expression":"2 @ 3"}`},
		{"MissingExpression", `{}`},
		{"UnknownField", `{"expression":"2+3","extra":1}`},
	}

	mux := newTestMux(t)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodPost,
				api.ExpressionEvaluate,
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

func TestExpressionAPIConfigurableLength(t *testing.T) {
	cfg := config.Config{
		MaxRequestBytes:     4096,
		MaxExpressionLength: 5,
	}

	mux := http.NewServeMux()
	NewHandler(cfg).RegisterRoutes(mux)

	req := httptest.NewRequest(
		http.MethodPost,
		api.ExpressionEvaluate,
		strings.NewReader(
			`{"expression":"100 + 200"}`,
		),
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
}

func TestExpressionAPIRequestLimit(t *testing.T) {
	cfg := config.Config{
		MaxRequestBytes:     10,
		MaxExpressionLength: 1024,
	}

	mux := http.NewServeMux()
	NewHandler(cfg).RegisterRoutes(mux)

	req := httptest.NewRequest(
		http.MethodPost,
		api.ExpressionEvaluate,
		strings.NewReader(
			`{"expression":"20 + 5"}`,
		),
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
}

func TestExpressionAPIMethodNotAllowed(t *testing.T) {
	mux := newTestMux(t)

	req := httptest.NewRequest(
		http.MethodGet,
		api.ExpressionEvaluate,
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
