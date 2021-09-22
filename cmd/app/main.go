package main

import (
	"flag"
	"net/http"

	"daitan-dispatch-system/cmd/app/routes"
)

func main() {

	// host := flag.String("--host", "localhost", "HTTP server network address")
	// port := flag.String("--port", "8080", "HTTP server network port")
	port := ":8080"

	flag.Parse()

	// info := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// error := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Client requests
	http.HandleFunc("/api/trips/", routes.RequestTrip)

	// Provider requests
	http.ListenAndServe(port, nil)
}
