package arithmetic

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"

	"calculator-app/backend/internal/api"
	"calculator-app/backend/internal/config"
)

type CalculateRequest struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	service *Service
	cfg     config.Config
}

// NewHandler uses dependency injection.
// Configuration is supplied by the application.
func NewHandler(service *Service, cfg config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(api.ArithmeticAdd, h.handleAdd)
	mux.HandleFunc(api.ArithmeticSubtract, h.handleSubtract)
	mux.HandleFunc(api.ArithmeticMultiply, h.handleMultiply)
	mux.HandleFunc(api.ArithmeticDivide, h.handleDivide)
}

func (h *Handler) handleAdd(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpAdd)
}

func (h *Handler) handleSubtract(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpSubtract)
}

func (h *Handler) handleMultiply(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpMultiply)
}

func (h *Handler) handleDivide(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpDivide)
}

func (h *Handler) handleOperation(
	w http.ResponseWriter,
	r *http.Request,
	operation api.Operation,
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

	// Request size is now controlled by config.Config.
	r.Body = http.MaxBytesReader(
		w,
		r.Body,
		h.cfg.MaxRequestBytes,
	)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request CalculateRequest

	if err := decoder.Decode(&request); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid JSON request",
		)
		return
	}

	// Reject additional JSON values after the request object.
	var extra interface{}
	if err := decoder.Decode(&extra); err != io.EOF {
		writeError(
			w,
			http.StatusBadRequest,
			"unexpected JSON content",
		)
		return
	}

	if request.A == nil || request.B == nil {
		writeError(
			w,
			http.StatusBadRequest,
			"both operands are required",
		)
		return
	}

	a := *request.A
	b := *request.B

	if !isFiniteOperand(a) || !isFiniteOperand(b) {
		writeError(
			w,
			http.StatusBadRequest,
			"operands must be finite numbers",
		)
		return
	}

	var result float64
	var err error

	switch operation {
	case api.OpAdd:
		result = h.service.Add(a, b)

	case api.OpSubtract:
		result = h.service.Subtract(a, b)

	case api.OpMultiply:
		result = h.service.Multiply(a, b)

	case api.OpDivide:
		result, err = h.service.Divide(a, b)

	default:
		err = errors.New("unsupported operation")
	}

	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if !isFiniteOperand(result) {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid calculation result",
		)
		return
	}

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(
		CalculateResponse{Result: result},
	)
}

func isFiniteOperand(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
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
