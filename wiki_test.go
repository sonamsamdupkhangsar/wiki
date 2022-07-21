package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLiveness(t *testing.T) {
	// Create a request to default path '/'
	req, err := http.NewRequest("GET", "/liveness", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(liveness)

	// Our handlers satisfy the http.Handler so can call serveHttp method
	handler.ServeHTTP(rr, req)

	// check the status code meets our expectation
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the reponse body is our expected
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("liveness handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestReadiness(t *testing.T) {
	req, err := http.NewRequest("GET", "/readiness", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(readiness)

	// Our handlers satisfy the http.Handler so can call serveHttp method
	handler.ServeHTTP(rr, req)

	// check the status code meets our expectation
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the reponse body is our expected
	expected := `{"ready": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
