package main

import (
	"log"
	"net/http"
	"time"

	"github.com/paisit04/shipping-go/config"
	"github.com/paisit04/shipping-go/handlers"
	"github.com/paisit04/shipping-go/handlers/rest"
	"github.com/paisit04/shipping-go/translation"
)

func main() {
	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := API(cfg)

	log.Printf("listening on %s\n", addr)

	srv := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           mux,
	}

	log.Fatal(srv.ListenAndServe())
}

func API(cfg config.Configuration) *http.ServeMux {
	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}
	if cfg.DatabaseURL != "" {
		db := translation.NewDatabaseService(cfg)
		translationService = db
	}

	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)
	return mux
}
