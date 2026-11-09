package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"suite-test/backend/internal/models"

	"github.com/jackc/pgx/v5"
)

func (h *StatsHandler) GetTopPublications(w http.ResponseWriter, r *http.Request) {
	// 1. Vérifier l'identifiant du client.
	clientID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || clientID <= 0 {
		sendError(w, http.StatusBadRequest, "Identifiant client invalide")
		return
	}

	// 2. Récupérer la limite demandée.
	limit := 5
	limitValue := r.URL.Query().Get("limit")

	if limitValue != "" {
		limit, err = strconv.Atoi(limitValue)
		if err != nil || limit < 1 || limit > 20 {
			sendError(w, http.StatusBadRequest, "La limite doit être comprise entre 1 et 20")
			return
		}
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
		sendError(w, http.StatusInternalServerError, "Erreur lors de la récupération du client")
		return
	}

	// 4. Récupérer les publications classées par engagement.
	rows, err := h.DB.Query(r.Context(), `
		SELECT
			id,
			platform,
			title,
			published_at,
			impressions,
			likes,
			comments,
			shares
		FROM publications
		WHERE client_id = $1
		ORDER BY
			CASE
				WHEN impressions = 0 THEN 0
				ELSE ROUND(
					(likes + comments + shares)::NUMERIC
					/ impressions * 100,
					2
				)
			END DESC,
			published_at DESC,
			id DESC
		LIMIT $2
	`, clientID, limit)

	if err != nil {
		sendError(w, http.StatusInternalServerError, "Erreur lors de la récupération des publications")
		return
	}
	defer rows.Close()

	publications := make([]models.PublicationTop, 0)

	for rows.Next() {
		var publication models.PublicationTop

		err := rows.Scan(
			&publication.ID,
			&publication.Platform,
			&publication.Title,
			&publication.PublishedAt,
			&publication.Impressions,
			&publication.Likes,
			&publication.Comments,
			&publication.Shares,
		)

		if err != nil {
			sendError(w, http.StatusInternalServerError, "Erreur lors de la lecture des publications")
			return
		}

		publication.Interactions = publication.Likes +
			publication.Comments +
			publication.Shares

		publication.EngagementRate = calculateEngagement(
			publication.Interactions,
			publication.Impressions,
		)

		publications = append(publications, publication)
	}

	if err := rows.Err(); err != nil {
		sendError(w, http.StatusInternalServerError, "Erreur lors de la lecture des publications")
		return
	}

	// 5. Retourner les publications en JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publications)
}
