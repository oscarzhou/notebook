package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetCat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(HandleGetCat))

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

	if string(b) != "Siamese" {
		t.Errorf("Expected body to be %q, but got %q", "Siamese", string(b))
	}
}
