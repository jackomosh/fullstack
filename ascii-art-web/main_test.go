package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestHomeHandler verifies that the home page loads correctly (200 OK)
// and that invalid paths return a 404.
func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
	}{
		{"Valid Path", "/", http.StatusOK},
		{"Invalid Path", "/something-else", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(homeHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					rr.Code, tt.expectedStatus)
			}
		})
	}
}

// TestAsciiHandler verifies the POST logic, including status codes for 
// bad requests and method limitations.
func TestAsciiHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		text           string
		banner         string
		expectedStatus int
	}{
		{"Valid Request", "POST", "Hello", "standard", http.StatusOK},
		{"Empty Text", "POST", "", "standard", http.StatusBadRequest},
		{"Invalid Banner", "POST", "Hello", "comic-sans", http.StatusBadRequest},
		{"Wrong Method (GET)", "GET", "Hello", "standard", http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create form data
			form := url.Values{}
			form.Add("text", tt.text)
			form.Add("banner", tt.banner)

			req, err := http.NewRequest(tt.method, "/ascii", strings.NewReader(form.Encode()))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(asciiHandler)

			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("%s: handler returned wrong status code: got %v want %v",
					tt.name, rr.Code, tt.expectedStatus)
			}
		})
	}
}