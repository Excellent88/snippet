package main

import (
	"fmt"
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
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal error server", 500)
	}
}
func showSnippet(w http.ResponseWriter, r *http.Request) { //snippet
	if r.URL.Path != "/snippet" {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, err = fmt.Fprintf(w, "id is: %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) { //snippet/create/
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		_, err := w.Write([]byte(`{"name":"Alex"}`))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Метод запрещен!", 405)
		return
	}
}
