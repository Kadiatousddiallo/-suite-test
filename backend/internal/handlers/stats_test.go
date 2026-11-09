package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"suite-test/backend/internal/database"
)

func TestCalculateEngagement(t *testing.T) {
	tests := []struct {
		name         string
		interactions int64
		impressions  int64
		expected     float64
	}{
		{
			name:         "calcul normal",
			interactions: 3040,
			impressions:  33700,
			expected:     9.02,
		},
		{
			name:         "aucune impression",
			interactions: 0,
			impressions:  0,
			expected:     0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculateEngagement(
				test.interactions,
				test.impressions,
			)

			if result != test.expected {
				t.Errorf(
					"Résultat attendu : %v, obtenu : %v",
					test.expected,
					result,
				)
			}
		})
	}
}

func TestGetClientStatsHTTP(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("Impossible de connecter la base de données : %v", err)
	}
	defer db.Close()

	handler := NewStatsHandler(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /clients/{id}/stats", handler.GetClientStats)

	t.Run("requête valide", func(t *testing.T) {
		request := httptest.NewRequest(
			http.MethodGet,
			"/clients/1/stats?from=2026-10-01&to=2026-10-31",
			nil,
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusOK {
			t.Errorf(
				"Statut attendu : 200, obtenu : %d",
				recorder.Code,
			)
		}

		if !strings.Contains(recorder.Body.String(), `"engagement_rate":9.02`) {
			t.Errorf("Le taux d'engagement attendu est absent de la réponse")
		}
	})

	t.Run("identifiant invalide", func(t *testing.T) {
		request := httptest.NewRequest(
			http.MethodGet,
			"/clients/abc/stats",
			nil,
		)

		recorder := httptest.NewRecorder()

		mux.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf(
				"Statut attendu : 400, obtenu : %d",
				recorder.Code,
			)
		}
	})
}
