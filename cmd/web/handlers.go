package main

import (
	"html/template"
	"log"
	"net/http"
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
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from showSnippet"))
}
func createSnippet(w http.ResponseWriter, r *http.Request) { //snippet/create/
	if r.Method != http.MethodPost {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{"name":"Alex"}`))
		http.Error(w, "Метод запрещен!", 405)
		return
	}
	w.Write([]byte(`{"name":"Alex"}`))
}
