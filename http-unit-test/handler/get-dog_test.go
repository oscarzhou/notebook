package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetDog_ByHostServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleGetDog))

	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if string(b) != "Corgi" {
		t.Errorf("Expected body to be %q, but got %q", "Corgi", string(b))
	}
}

func TestHandleGetDog_BySimRequest(t *testing.T) {
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
