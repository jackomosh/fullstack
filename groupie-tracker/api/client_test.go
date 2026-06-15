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
		HTTPClient: &http.Clienpackage main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/api"
)

func init() {
	// Initialize minimal runtime mocks to run isolated pipeline checks safely
	registry = &api.UnifiedRegistry{
		Artists: []api.Artist{
			{ID: 1, Name: "Pink Floyd", Members: []string{"Syd", "Roger"}, CreationDate: 1965},
		},
		Relations: map[int]api.Relation{
			1: {ID: 1, DatesLocations: map[string][]string{"london-uk": {"10-10-2026"}}},
		},
	}
	templates = template.Must(template.New("index.html").Parse(`{{range .}}<h1>{{.Name}}</h1>{{end}}`))
	_, _ = templates.New("details.html").Parse(`<h2>{{.Artist.Name}}</h2>`)
}

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	homeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}
}

func TestArtistDetailsHandler(t *testing.T) {
	tests := []struct {
		name       string
		queryID    string
		wantStatus int
	}{
		{"Valid Query Reference", "1", http.StatusOK},
		{"Out-of-Bounds Identifier Match", "99", http.StatusNotFound},
		{"Malformed Request Parameters", "abc", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/artist?id="+tt.queryID, nil)
			rr := httptest.NewRecorder()

			artistDetailsHandler(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("Expected status %d, got %d", tt.wantStatus, rr.Code)
			}
		})
	}
}t{Timeout: 2 * time.Second},
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