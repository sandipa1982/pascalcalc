package expression

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"unicode/utf8"

	"calculator-app/backend/internal/api"
	"calculator-app/backend/internal/config"
)

type EvaluateRequest struct {
	Expression string `json:"expression"`
}

type EvaluateResponse struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	cfg config.Config
}

// NewHandler receives configuration from the application.
func NewHandler(cfg config.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(
		api.ExpressionEvaluate,
		h.handleEvaluate,
	)
}

func (h *Handler) handleEvaluate(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeError(
			w,
			http.StatusMethodNotAllowed,
			"method not allowed",
		)
		return
	}

	// Apply configured HTTP request-size limit.
	r.Body = http.MaxBytesReader(
		w,
		r.Body,
		h.cfg.MaxRequestBytes,
	)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request EvaluateRequest

	if err := decoder.Decode(&request); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid JSON request",
		)
		return
	}

	// Reject additional JSON values.
	var extra interface{}
	if err := decoder.Decode(&extra); err != io.EOF {
		writeError(
			w,
			http.StatusBadRequest,
			"unexpected JSON content",
		)
		return
	}

	expr := strings.TrimSpace(request.Expression)

	if expr == "" {
		writeError(
			w,
			http.StatusBadRequest,
			"expression is required",
		)
		return
	}

	if !utf8.ValidString(expr) {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid expression encoding",
		)
		return
	}

	// Create a fresh evaluator for each request.
	// Evaluator contains mutable parsing state.
	evaluator := NewEvaluatorWithLimit(
		h.cfg.MaxExpressionLength,
	)

	result, err := evaluator.Evaluate(expr)
	if err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(
		EvaluateResponse{
			Expression: expr,
			Result:     result,
		},
	)
}

func writeError(
	w http.ResponseWriter,
	status int,
	message string,
) {
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(
		ErrorResponse{Error: message},
	)
}
