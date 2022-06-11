package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		htmlFolder + "home.page.tmpl",
		htmlFolder + "base.layout.tmpl",
		htmlFolder + "footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Cannot process the template for the home page: %v\n", err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Printf("Cannot execute the template of the home page: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// w.Write
}
