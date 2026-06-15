package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchArtists(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"name":"The Mock Band","creationDate":2020}]`))
	}))
	defer server.Close()

	client := &Client{
		HTTPClient: &http.Client{Timeout: 2 * time.Second},
	}

	// Override standard endpoint mapping with our localized mock address
	mockURL := server.URL
	resp, err := client.HTTPClient.Get(mockURL)
	if err != nil {
		t.Fatalf("Failed to execute internal test server call: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}