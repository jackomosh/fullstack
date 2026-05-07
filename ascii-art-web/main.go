package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"ascii-art-web/ascii"
)

// PageData holds the information sent to the HTML template
type PageData struct {
	Result string
}

func main() {
	// GET route for the home page
	http.HandleFunc("/", homeHandler)
	// POST route for processing arttemplate
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	renderTemplate(w, "index.html", nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    text := r.FormValue("text")
    banner := r.FormValue("banner")

    // Validation for empty text or invalid banners
    if text == "" || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {
        http.Error(w, "400 Bad Request", http.StatusBadRequest)
        return
    }

    // Call our new library function
    result, err := ascii.GenerateAscii(text, banner)
    if err != nil {
        // If the file is missing, return 404, else 500
        if os.IsNotExist(err) {
            http.Error(w, "404 Banner Not Found", http.StatusNotFound)
        } else {
            http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
        }
        return
    }

    renderTemplate(w, "index.html", PageData{Result: result})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}