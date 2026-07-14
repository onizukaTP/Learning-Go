package main

import (
	"log"
	"net/http"
)

// Define a home handler func which writes a byte slice containing
// "Hello from snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Use the http.NewServeMux() func to initialize a new servemux, then
	// register the home func as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() func to start a new web server. We pass in
	// two params: the TCP network address to listen on
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() func to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
