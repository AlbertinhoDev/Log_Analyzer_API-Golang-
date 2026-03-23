package main

import (
	"log"
	"loganalyzerapi/internal/handler"
	"loganalyzerapi/internal/service"
	"net/http"
	"time"
)

func main() {
	parser := service.NewParser()

	analyzer := service.NewAnalyzer(parser)

	logsHandler := handler.NewLogsHandler(analyzer)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", logsHandler.Health)

	mux.HandleFunc("/analyze", logsHandler.Analyze)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Println("server started on :8080")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
