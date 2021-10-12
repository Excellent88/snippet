package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	fmt.Println("Allo")
	fmt.Println("Allo")
	fmt.Println("Allo")
	fmt.Println("Allo")
	fmt.Println("Allo")
	fmt.Println("Allo")
	//flag
	addr := flag.String("addr", ":4000", "Сетевой адрес веб-сервера")
	flag.Parse()
	//Logger
	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
	//Handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create/", app.createSnippet)
	//Files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//Server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Printf("Запуск сервера на %s", *addr)
	err = srv.ListenAndServe() //addr = "host:port"
	errorLog.Fatal(err)
}
