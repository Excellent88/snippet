package main

import (
	"fmt"
	"golangify.com/snippetbox/morestrings"
	"golangify.com/snippetbox/newpkg"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
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

func main() {
	newpkg.NewMessage()
	fmt.Println(morestrings.ReverseRunes("ABCDEFG"))

	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create/", createSnippet)

	fmt.Println("Запуск сервера на порту :4000")
	err := http.ListenAndServe(":4000", mux) //addr = "host:port"
	log.Fatal(err)
}
