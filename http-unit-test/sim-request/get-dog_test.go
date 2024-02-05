package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetDog(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	HandleGetDog(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "Corgi" {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "Corgi")
	}
}
