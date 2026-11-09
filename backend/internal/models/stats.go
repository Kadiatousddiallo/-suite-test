package models

type ClientStats struct {
	ClientID   int             `json:"client_id"`
	ClientName string          `json:"client_name"`
	Total      StatsTotals     `json:"total"`
	ByPlatform []PlatformStats `json:"by_platform"`
}

type StatsTotals struct {
	Publications   int     `json:"publications"`
	Impressions    int64   `json:"impressions"`
	Interactions   int64   `json:"interactions"`
	EngagementRate float64 `json:"engagement_rate"`
}

type PlatformStats struct {
	Platform       string  `json:"platform"`
	Publications   int     `json:"publications"`
	Impressions    int64   `json:"impressions"`
	Interactions   int64   `json:"interactions"`
	EngagementRate float64 `json:"engagement_rate"`
}
