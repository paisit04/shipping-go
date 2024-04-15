package main

import (
	"log"
	"net/http"

	"github.com/paisit04/shipping-go/handlers/rest"
)

func main() {
	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
