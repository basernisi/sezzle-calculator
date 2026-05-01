package httpadapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/basernisi/sezzle-calculator/backend/internal/application/calculate"
	"github.com/basernisi/sezzle-calculator/backend/internal/domain/calculator"
)

const maxRequestBodyBytes int64 = 1024

type CalculateService interface {
	Execute(request calculate.Request) (calculate.Response, error)
}

type Handler struct {
	service CalculateService
}

func NewHandler(service CalculateService) Handler {
	return Handler{service: service}
}

func (h Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/calculate", h.Calculate)
}

func (h Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodyBytes)
	defer r.Body.Close()

	var requestBody CalculateRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&requestBody); err != nil {
		h.handleDecodeError(w, err)
		return
	}

	var extraPayload struct{}
	if err := decoder.Decode(&extraPayload); err != io.EOF {
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "Request body must contain a single JSON object")
		return
	}

	response, err := h.service.Execute(calculate.Request{
		Operation: requestBody.Operation,
		Operands:  requestBody.Operands,
	})
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, CalculateResponse{Result: response.Result})
}

func (h Handler) handleDecodeError(w http.ResponseWriter, err error) {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxError):
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "Malformed JSON payload")
	case errors.As(err, &unmarshalTypeError):
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Request contains invalid field types")
	case errors.Is(err, io.EOF):
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "Request body is required")
	case err.Error() == "http: request body too large":
		writeError(w, http.StatusRequestEntityTooLarge, "REQUEST_TOO_LARGE", "Request body exceeds size limit")
	default:
		writeError(w, http.StatusBadRequest, "INVALID_JSON", "Request body is invalid")
	}
}

func (h Handler) handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, calculate.ErrInvalidRequest):
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Operation and operands must be valid finite numbers")
	case errors.Is(err, calculator.ErrUnsupportedOperation):
		writeError(w, http.StatusBadRequest, "UNSUPPORTED_OPERATION", "The requested operation is not supported")
	case errors.Is(err, calculator.ErrInvalidOperandCount):
		writeError(w, http.StatusBadRequest, "INVALID_OPERAND_COUNT", "The provided operands do not match the operation requirements")
	case errors.Is(err, calculator.ErrDivisionByZero):
		writeError(w, http.StatusBadRequest, "DIVISION_BY_ZERO", "Division by zero is not allowed")
	case errors.Is(err, calculator.ErrNegativeSquareRoot):
		writeError(w, http.StatusBadRequest, "NEGATIVE_SQRT", "Square root of a negative number is not allowed")
	default:
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected error occurred")
	}
}
