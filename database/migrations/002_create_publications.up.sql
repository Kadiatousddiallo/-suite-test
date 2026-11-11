
CREATE TABLE IF NOT EXISTS publications (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES clients(id),
    platform TEXT NOT NULL CHECK (
        platform IN ('instagram', 'facebook', 'tiktok', 'linkedin')
    ),
    title TEXT NOT NULL,
    published_at TIMESTAMPTZ NOT NULL,
    impressions INTEGER NOT NULL DEFAULT 0,
    likes INTEGER NOT NULL DEFAULT 0,
    comments INTEGER NOT NULL DEFAULT 0,
    shares INTEGER NOT NULL DEFAULT 0
);
