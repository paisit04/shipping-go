package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/paisit04/shipping-go/handlers"
	"github.com/paisit04/shipping-go/handlers/rest"
	"github.com/paisit04/shipping-go/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)

	srv := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	log.Fatal(srv.ListenAndServe())
}
