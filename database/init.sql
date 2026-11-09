CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE publications (
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

INSERT INTO clients (name) VALUES
    ('Boulangerie Kaloum'),
    ('Studio Dakar Mode');

INSERT INTO publications (
    client_id, platform, title, published_at,
    impressions, likes, comments, shares
) VALUES
    (1, 'instagram', 'Nouveau pain au fonio',
     '2026-10-02 09:00+00', 4200, 310, 45, 22),

    (1, 'instagram', 'Coulisses du fournil',
     '2026-10-09 18:30+00', 3800, 290, 30, 15),

    (1, 'facebook', 'Horaires de la Toussaint',
     '2026-10-12 08:00+00', 2500, 80, 12, 9),

    (1, 'tiktok', 'Recette express en 30 s',
     '2026-10-15 19:00+00', 12500, 980, 120, 210),

    (1, 'linkedin', 'Nous recrutons un boulanger',
     '2026-10-20 07:45+00', 1600, 40, 8, 14),

    (1, 'tiktok', 'Test de goût en famille',
     '2026-10-25 20:00+00', 9100, 640, 85, 130),

    (1, 'facebook', 'Publication programmée',
     '2026-10-31 10:00+00', 0, 0, 0, 0),

    (2, 'instagram', 'Collection automne',
     '2026-10-05 12:00+00', 6100, 520, 60, 40),

    (2, 'tiktok', 'Essayage en direct',
     '2026-10-18 17:00+00', 15400, 1200, 150, 300);
