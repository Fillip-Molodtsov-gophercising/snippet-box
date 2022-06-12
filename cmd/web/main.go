package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const (
	htmlFolder   = "./ui/html/"
	staticFolder = "./ui/static/"
)

func main() {
	addr := flag.String("addr", ":4040", "Address of the server")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		errorLog.Println(err)
	}
}
