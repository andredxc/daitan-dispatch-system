package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	error   *log.Logger
	info    *log.Logger
	service *Service
}

func main() {

	host := flag.String("--host", "localhost", "HTTP server network address")
	port := flag.Int("--port", 8080, "HTTP server network port")

	flag.Parse()

	info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	error := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		info:    info,
		error:   error,
		service: new(Service),
	}

	serverURI := fmt.Sprintf("%s:%d", *host, *port)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     error,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	info.Printf("Starting server on %s", serverURI)
	err := srv.ListenAndServe()
	error.Fatal(err)
}
