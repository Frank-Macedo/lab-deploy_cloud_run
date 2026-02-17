package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTemperature_InvalidCep(t *testing.T) {
	req := httptest.NewRequest("GET", "/temperature/invalid", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/temperature/{cep}", GetTemperature)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status %d, got %d", http.StatusUnprocessableEntity, w.Code)
	}
	if w.Body.String() != "invalid zipcode\n" {
		t.Errorf("expected 'invalid zipcode', got %s", w.Body.String())
	}
}

func TestGetTemperature_ValidCep_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/temperature/01310100", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/temperature/{cep}", GetTemperature)
	router.ServeHTTP(w, req)

	if w.Code == http.StatusUnprocessableEntity {
		t.Errorf("valid CEP returned error: %s", w.Body.String())
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetTemperature_MissingCep(t *testing.T) {
	req := httptest.NewRequest("GET", "/temperature/", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/temperature/{cep}", GetTemperature)
	router.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		t.Error("expected error for missing CEP")
	}
}
