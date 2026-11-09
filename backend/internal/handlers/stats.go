package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"time"

	"suite-test/backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StatsHandler struct {
	DB *pgxpool.Pool
}

func NewStatsHandler(db *pgxpool.Pool) *StatsHandler {
	return &StatsHandler{DB: db}
}

func sendError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func calculateEngagement(interactions, impressions int64) float64 {
	if impressions == 0 {
		return 0
	}

	rate := float64(interactions) / float64(impressions) * 100

	return math.Round(rate*100) / 100
}

func (h *StatsHandler) GetClientStats(w http.ResponseWriter, r *http.Request) {
	// 1. Récupérer et vérifier l'identifiant du client.
	clientID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || clientID <= 0 {
		sendError(w, http.StatusBadRequest, "Identifiant client invalide")
		return
	}

	// 2. Récupérer les dates facultatives.
	fromValue := r.URL.Query().Get("from")
	toValue := r.URL.Query().Get("to")

	var fromDate, toDate *time.Time

	if fromValue != "" {
		parsed, err := time.Parse("2006-01-02", fromValue)
		if err != nil {
			sendError(w, http.StatusBadRequest, "Date de début invalide")
			return
		}
		fromDate = &parsed
	}

	if toValue != "" {
		parsed, err := time.Parse("2006-01-02", toValue)
		if err != nil {
			sendError(w, http.StatusBadRequest, "Date de fin invalide")
			return
		}
		toDate = &parsed
	}

	if fromDate != nil && toDate != nil && fromDate.After(*toDate) {
		sendError(w, http.StatusBadRequest, "La date de début doit précéder la date de fin")
		return
	}

	// 3. Vérifier que le client existe.
	var clientName string

	err = h.DB.QueryRow(
		r.Context(),
		"SELECT name FROM clients WHERE id = $1",
		clientID,
	).Scan(&clientName)

	if err == pgx.ErrNoRows {
		sendError(w, http.StatusNotFound, "Client introuvable")
		return
	}

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Erreur lors de la récupération des statistiques")
		return
	}

	// 4. Calculer les statistiques générales.
	var total models.StatsTotals

	err = h.DB.QueryRow(r.Context(), `
		SELECT
			COUNT(*),
			COALESCE(SUM(impressions), 0),
			COALESCE(SUM(likes + comments + shares), 0)
		FROM publications
		WHERE client_id = $1
		AND ($2::date IS NULL OR published_at >= $2::date)
		AND ($3::date IS NULL OR published_at < $3::date + INTERVAL '1 day')
	`, clientID, fromDate, toDate).Scan(
		&total.Publications,
		&total.Impressions,
		&total.Interactions,
	)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Erreur lors du calcul des statistiques")
		return
	}

	total.EngagementRate = calculateEngagement(
		total.Interactions,
		total.Impressions,
	)

	// 5. Calculer les statistiques par plateforme.
	rows, err := h.DB.Query(r.Context(), `
		SELECT
			platform,
			COUNT(*),
			COALESCE(SUM(impressions), 0),
			COALESCE(SUM(likes + comments + shares), 0)
		FROM publications
		WHERE client_id = $1
		AND ($2::date IS NULL OR published_at >= $2::date)
		AND ($3::date IS NULL OR published_at < $3::date + INTERVAL '1 day')
		GROUP BY platform
		ORDER BY platform
	`, clientID, fromDate, toDate)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Erreur lors du calcul par plateforme")
		return
	}
	defer rows.Close()

	platforms := make([]models.PlatformStats, 0)

	for rows.Next() {
		var platform models.PlatformStats

		err := rows.Scan(
			&platform.Platform,
			&platform.Publications,
			&platform.Impressions,
			&platform.Interactions,
		)

		if err != nil {
			sendError(w, http.StatusInternalServerError, "Erreur lors de la lecture des statistiques")
			return
		}

		platform.EngagementRate = calculateEngagement(
			platform.Interactions,
			platform.Impressions,
		)

		platforms = append(platforms, platform)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, "Erreur lors de la lecture des statistiques")
		return
	}

	// 6. Construire et retourner la réponse JSON.
	response := models.ClientStats{
		ClientID:   clientID,
		ClientName: clientName,
		Total:      total,
		ByPlatform: platforms,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
