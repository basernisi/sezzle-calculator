package httpadapter

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/basernisi/sezzle-calculator/backend/internal/application/calculate"
	"github.com/basernisi/sezzle-calculator/backend/internal/domain/calculator"
)

func TestHandlerCalculateSuccess(t *testing.T) {
	service := calculate.NewService(calculate.NewOperationRegistry(calculator.AddOperation{}))
	handler := NewHandler(service)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(`{"operation":"add","operands":[10,5]}`))
	responseRecorder := httptest.NewRecorder()

	handler.Calculate(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}
}

func TestHandlerCalculateDivisionByZero(t *testing.T) {
	service := calculate.NewService(calculate.NewOperationRegistry(calculator.DivideOperation{}))
	handler := NewHandler(service)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(`{"operation":"divide","operands":[10,0]}`))
	responseRecorder := httptest.NewRecorder()

	handler.Calculate(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestHandlerCalculateInvalidJSON(t *testing.T) {
	service := calculate.NewService(calculate.NewOperationRegistry(calculator.AddOperation{}))
	handler := NewHandler(service)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(`{"operation":"add","operands":[10,5]`))
	responseRecorder := httptest.NewRecorder()

	handler.Calculate(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}

func TestHandlerCalculateRejectsTrailingJSON(t *testing.T) {
	service := calculate.NewService(calculate.NewOperationRegistry(calculator.AddOperation{}))
	handler := NewHandler(service)

	request := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBufferString(`{"operation":"add","operands":[10,5]}{"extra":true}`))
	responseRecorder := httptest.NewRecorder()

	handler.Calculate(responseRecorder, request)

	if responseRecorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, responseRecorder.Code)
	}
}
