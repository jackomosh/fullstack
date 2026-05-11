package main

import (
	"net/http"
	"fmt"
	"html/template"
	"os"
)

type PageData struct {
	Result string
}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)

	fmt.Println("Starting server at http://localhost:5000")
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Printf("Server failed %s\n", err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println(w, "404 Not Found", http.StatusNotFound)
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

	if text == "" || (banner != "standard" && banner != "shadow" && banner != "thinkertoy") {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	result, err := ascii.GenerateAscii(text, banner)

	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404 Not Found", http.StatusNotFound)
		} else {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	renderTemplate(w, "index.html", PageData{Result: result})

}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("/templates", tmpl)
	if err != nil {
		http.Error(w, "500 internal server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}