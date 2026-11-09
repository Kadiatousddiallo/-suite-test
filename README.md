\# SUITE — Module de statistiques des publications



\## Présentation



Application web de suivi des statistiques des publications sur les réseaux sociaux pour les agences marketing.



Le projet permet de consulter les performances des publications par client et par plateforme : Instagram, Facebook, TikTok et LinkedIn.



\## Technologies utilisées



\### Backend

\- Go

\- API REST avec `net/http`

\- PostgreSQL

\- pgx



\### Frontend

\- Next.js

\- React

\- TypeScript

\- Tailwind CSS



\### Infrastructure

\- Docker

\- Docker Compose



\## Fonctionnalités



\- Consultation des statistiques d'un client.

\- Filtrage des statistiques par période.

\- Répartition des performances par plateforme.

\- Calcul du taux d'engagement.

\- Classement des meilleures publications.

\- Gestion des erreurs et des états de chargement.



\## Structure du projet



```text

suite-test/

├── backend/

│   ├── internal/

│   │   ├── database/

│   │   ├── handlers/

│   │   └── models/

│   ├── go.mod

│   └── main.go

├── database/

│   └── init.sql

├── frontend/

│   └── src/

│       └── app/

├── docker-compose.yml

└── README.md

```



\## Installation et lancement



\### 1. Démarrer PostgreSQL



À la racine du projet :



```bash

docker compose up -d

```



\### 2. Démarrer le backend Go



Dans un autre terminal :



```bash

cd backend

go mod tidy

go run .

```



L'API est accessible sur :



`http://localhost:8080`



\### 3. Démarrer le frontend



Dans un autre terminal :



```bash

cd frontend

npm install

npm run dev

```



L'application est accessible sur :



`http://localhost:3000`



Exemple de tableau de bord :



`http://localhost:3000/clients/1`



\## Endpoints API



| Méthode | Endpoint | Description |

|---|---|---|

| GET | `/clients/{id}/stats` | Statistiques d'un client |

| GET | `/clients/{id}/publications/top` | Meilleures publications |



Exemple :



`GET /clients/1/stats?from=2026-10-01\&to=2026-10-31`



\## Tests



Pour exécuter les tests du backend :



```bash

cd backend

go test ./... -v

```



\## Choix techniques



\- Go avec `net/http` pour développer une API REST légère.

\- PostgreSQL pour stocker les clients et les publications.

\- pgx pour la connexion à PostgreSQL.

\- Next.js et TypeScript pour construire l'interface.

\- Docker Compose pour faciliter le démarrage de la base de données.



\## Améliorations possibles



\- Ajouter de nouveaux indicateurs statistiques.

\- Améliorer la visualisation des données.

\- Ajouter des fonctionnalités de reporting.



\## Temps consacré



À compléter avec le temps réellement consacré au test.



\## Auteur



Kadiatou Sadio Diallo

