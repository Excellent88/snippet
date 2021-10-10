package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create/", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	flag.Parse()

	log.Printf("Connecting address is: %s", *addr)
	err := http.ListenAndServe(*addr, mux) //addr = "host:port"
	log.Fatal(err)
}
