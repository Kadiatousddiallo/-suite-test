package models

import "time"

type PublicationTop struct {
	ID             int       `json:"id"`
	Platform       string    `json:"platform"`
	Title          string    `json:"title"`
	PublishedAt    time.Time `json:"published_at"`
	Impressions    int64     `json:"impressions"`
	Likes          int64     `json:"likes"`
	Comments       int64     `json:"comments"`
	Shares         int64     `json:"shares"`
	Interactions   int64     `json:"interactions"`
	EngagementRate float64   `json:"engagement_rate"`
}
