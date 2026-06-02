package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// PageData allows us to pass dynamic content to the professional UI
type PageData struct {
	Title string
	Year  int
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := PageData{
		Title: "IFC | Integrity Framework for Capital",
		Year:  time.Now().Year(),
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Template Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}

func main() {
	// Serve static files (CSS, Images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", homeHandler)

	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}