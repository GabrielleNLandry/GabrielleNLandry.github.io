package main	

import (
	"html/template"
	"log"
	"net/http"
)

func hobbyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/hobby.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
	http.Error(w, "Error loading template", http.StatusInternalServerError)
	return
	}
	tmpl.Execute(w, nil)
	}

//Handler Function for the Homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homeHandler) // Route for the homepage
	http.HandleFunc("/about", aboutHandler) // Route for the about page
	http.HandleFunc("/hobby", hobbyHandler) // Route for the hobby page

	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Start the server on port 8080
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}