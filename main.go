package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("Starting skool")
	http.Handle("/public/", http.StripPrefix(strings.TrimRight("/public/", "/"), http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("indexHandler called")
	tmpl := template.New("indexTemplate")
	tmpl, err := tmpl.ParseFiles("templates/layouts/default.html", "templates/index.html")
	if err != nil {
		log.Printf("Error parsing the files: %v", err)
		return
	}

	// err = tmpl.Execute(w, nil)
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error executing the template: %v", err)
		return
	}
	// tmpl.ExecuteTemplate(w, "index", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("aboutHandler called")
	tmpl := template.New("indexTemplate")
	tmpl, err := tmpl.ParseFiles("templates/layouts/default.html", "templates/about.html")
	if err != nil {
		log.Printf("Error parsing the files: %v", err)
		return
	}

	// err = tmpl.Execute(w, nil)
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("Error executing the template: %v", err)
		return
	}

}
