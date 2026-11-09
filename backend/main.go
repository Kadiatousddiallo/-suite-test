package main

import (
	"log"
	"net/http"

	"suite-test/backend/internal/database"
	"suite-test/backend/internal/handlers"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Erreur de connexion à PostgreSQL :", err)
	}
	defer db.Close()

	statsHandler := handlers.NewStatsHandler(db)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /clients/{id}/stats", statsHandler.GetClientStats)
	mux.HandleFunc("GET /clients/{id}/publications/top", statsHandler.GetTopPublications)

	log.Println("API SUITE démarrée sur http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Erreur du serveur HTTP :", err)
	}
}
