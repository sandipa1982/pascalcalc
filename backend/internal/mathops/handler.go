package mathops

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"calculator-app/backend/internal/api"
	"calculator-app/backend/internal/config"
)

type MathRequest struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type MathResponse struct {
	Result float64 `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	service *Service
	cfg     config.Config
}

// NewHandler injects the business service and configuration.
func NewHandler(service *Service, cfg config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc(api.MathPower, h.handlePower)
	mux.HandleFunc(api.MathSquareRoot, h.handleSquareRoot)
	mux.HandleFunc(api.MathPercentage, h.handlePercentage)
}

func (h *Handler) handlePower(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpPower)
}

func (h *Handler) handleSquareRoot(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpSquareRoot)
}

func (h *Handler) handlePercentage(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.handleOperation(w, r, api.OpPercentage)
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

	// Read request size from configuration.
	r.Body = http.MaxBytesReader(
		w,
		r.Body,
		h.cfg.MaxRequestBytes,
	)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var request MathRequest

	if err := decoder.Decode(&request); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"invalid JSON request",
		)
		return
	}

	// Reject multiple JSON values in the same request body.
	var extra interface{}
	if err := decoder.Decode(&extra); err != io.EOF {
		writeError(
			w,
			http.StatusBadRequest,
			"unexpected JSON content",
		)
		return
	}

	if request.A == nil {
		writeError(
			w,
			http.StatusBadRequest,
			"operand a is required",
		)
		return
	}

	if operation != api.OpSquareRoot && request.B == nil {
		writeError(
			w,
			http.StatusBadRequest,
			"operand b is required",
		)
		return
	}

	var result float64
	var err error

	switch operation {
	case api.OpPower:
		result, err = h.service.Power(
			*request.A,
			*request.B,
		)

	case api.OpSquareRoot:
		result, err = h.service.SquareRoot(
			*request.A,
		)

	case api.OpPercentage:
		result, err = h.service.Percentage(
			*request.A,
			*request.B,
		)

	default:
		err = errors.New("unsupported operation")
	}

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
		MathResponse{Result: result},
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
