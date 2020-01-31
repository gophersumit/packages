package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("fresh package : hot reloading for your apps!!!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRequest)
	fmt.Printf("starting server on %d\n", 3000)
	http.ListenAndServe(":3000", mux)
}
