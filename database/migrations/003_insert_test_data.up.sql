
-- Insérer les clients avec leurs identifiants d'origine
INSERT INTO clients (id, name) VALUES
    (1, 'Boulangerie Kaloum'),
    (2, 'Studio Dakar Mode')
ON CONFLICT (id) DO NOTHING;

-- Insérer les publications
INSERT INTO publications (
    id, client_id, platform, title, published_at,
    impressions, likes, comments, shares
) VALUES
    (1, 1, 'instagram', 'Nouveau pain au fonio',
     '2026-10-02 09:00+00', 4200, 310, 45, 22),

    (2, 1, 'instagram', 'Coulisses du fournil',
     '2026-10-09 18:30+00', 3800, 290, 30, 15),

    (3, 1, 'facebook', 'Horaires de la Toussaint',
     '2026-10-12 08:00+00', 2500, 80, 12, 9),

    (4, 1, 'tiktok', 'Recette express en 30 s',
     '2026-10-15 19:00+00', 12500, 980, 120, 210),

    (5, 1, 'linkedin', 'Nous recrutons un boulanger',
     '2026-10-20 07:45+00', 1600, 40, 8, 14),

    (6, 1, 'tiktok', 'Test de goût en famille',
     '2026-10-25 20:00+00', 9100, 640, 85, 130),

    (7, 1, 'facebook', 'Publication programmée',
     '2026-10-31 10:00+00', 0, 0, 0, 0),

    (8, 2, 'instagram', 'Collection automne',
     '2026-10-05 12:00+00', 6100, 520, 60, 40),

    (9, 2, 'tiktok', 'Essayage en direct',
     '2026-10-18 17:00+00', 15400, 1200, 150, 300)
ON CONFLICT (id) DO NOTHING;

-- Synchroniser les séquences SERIAL avec les identifiants existants
SELECT setval(
    pg_get_serial_sequence('clients', 'id'),
    (SELECT MAX(id) FROM clients),
    true
);

SELECT setval(
    pg_get_serial_sequence('publications', 'id'),
    (SELECT MAX(id) FROM publications),
    true
);
