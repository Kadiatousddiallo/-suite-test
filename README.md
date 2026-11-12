SUITE – Module de statistiques des publications

# 1. Présentation
Ce projet a été réalisé dans le cadre du test technique SUITE.

Il permet de consulter les statistiques des publications d'un client sur quatre réseaux sociaux : Instagram, Facebook, TikTok et LinkedIn.

L'application affiche les impressions, les interactions, le taux d'engagement et les meilleures publications.

# 2. Technologies et choix techniques

- **Go et net/http :** utilisés pour développer l'API REST. J'ai choisi Go pour sa simplicité et ses performances, et `net/http` pour éviter d'ajouter un framework inutile.
- **PostgreSQL et pgx :** PostgreSQL permet de stocker les clients et leurs publications. `pgx` assure la communication entre l'API Go et la base de données.
- **Next.js et TypeScript :** utilisés pour développer le frontend. TypeScript permet de mieux structurer le code et de limiter les erreurs de typage.
- **Chart.js :** utilisé pour représenter les statistiques par plateforme sous forme de graphiques.
- **Docker Compose :** permet de lancer PostgreSQL, les migrations, l'API et le frontend avec une seule commande.
- **golang-migrate :** utilisé pour gérer les migrations SQL et organiser la création et l'évolution de la base de données.


# 3. Lancer le projet en local
Le projet peut être lancé de deux manières : d'abord sans Docker, en démarrant chaque composant séparément, puis avec Docker Compose, qui automatise le démarrage de l'ensemble.
## A. Avant la dockerisation : lancement séparé des composants
Prérequis :
- Go installé.
- Node.js et npm installés.
- PostgreSQL installé et démarré.
Étape 1 – Préparer la base de données

Créer l'utilisateur et la base de données PostgreSQL :
CREATE USER suite_user WITH PASSWORD 'suite_password';
CREATE DATABASE suite_db OWNER suite_user;
Depuis la racine du projet, exécuter le script SQL pour créer les tables et insérer les données de test :
psql -U suite_user -d suite_db -f database/init.sql
Étape 2 – Installer et lancer l'API Go

Dans un premier terminal :
cd backend
go mod download
Configurer la connexion à PostgreSQL (dans Git Bash) :
export DATABASE_URL="postgres://suite_user:suite_password@localhost:5432/suite_db?sslmode=disable"
Exécuter les tests du backend :
go test ./... -v
Démarrer l'API :
go run .
L'API est accessible sur http://localhost:8080.

Étape 3 – Installer et lancer le frontend

Dans un deuxième terminal :
cd frontend
npm install
npm run dev
Le frontend est accessible sur http://localhost:3000.

Pour consulter le tableau de bord :

http://localhost:3000/clients/1

## B. Après la dockerisation : lancement avec Docker Compose
J'ai ensuite utilisé Docker pour faciliter le lancement du projet et éviter de devoir démarrer chaque service manuellement.

Docker Compose permet de démarrer PostgreSQL, les migrations, l'API et le frontend.

Étape 1 – Construire les images et démarrer les services

À la racine du projet :
docker compose up -d --build
Cette commande construit les images Docker et démarre les différents conteneurs.

Étape 2 – Vérifier l'état des services
docker compose ps -a
Étape 3 – Vérifier l'exécution des migrations
docker compose logs migrate
Les migrations créent automatiquement les tables et insèrent les données de test avant le démarrage de l'API.

Étape 4 – Consulter les logs de l'API
docker compose logs api
Étape 5 – Accéder à l'application
- Frontend : http://localhost:3000/clients/1
- API : http://localhost:8080
- PostgreSQL : localhost:5439
Étape 6 – Arrêter les services
docker compose down
Cette commande arrête et supprime les conteneurs, mais conserve le volume de données PostgreSQL.

Attention : ne pas utiliser docker compose down -v si l'on souhaite conserver les données de la base.

# 4. Endpoints de l'API
Méthode	Endpoint	Description
GET	/clients/{id}/stats	Récupérer les statistiques d'un client
GET	/clients/{id}/publications/top	Récupérer les meilleures publicationsExemples de requêtes :
curl "http://localhost:8080/clients/1/stats?from=2026-10-01&to=2026-10-31"

curl "http://localhost:8080/clients/1/publications/top?limit=3"
Le filtre par dates est facultatif. La limite des meilleures publications est de 5 par défaut et ne peut pas dépasser 20.

# 5. Tests réalisés
Pour lancer les tests du backend :
cd backend
go test ./... -v
Les tests vérifient notamment :
- Le calcul du taux d'engagement.
- Le cas où les impressions sont nulles.
- Le fonctionnement du handler HTTP.
- Le rejet d'un identifiant invalide.
J'ai également effectué des tests manuels des endpoints avec curl.


# 6. Fonctionnalité non réalisée et améliorations possibles
Je n'ai pas eu le temps de réaliser le bonus de génération automatique de rapports avec une intelligence artificielle.

Pour l'ajouter, j'aurais développé un endpoint Go chargé de récupérer les statistiques d'un client, de préparer un résumé des indicateurs et d'envoyer ces informations à une API d'intelligence artificielle.

La clé API aurait été stockée dans une variable d'environnement afin de ne pas l'exposer dans le code.

J'aurais également ajouté davantage de tests pour couvrir les différents cas d'erreur.

# 7. Temps réellement consacré
Environ 5 jours.

Ce temps comprend le développement, les recherches et l'apprentissage des bases de Go à travers la documentation officielle, des tutoriels et des outils d'assistance comme Claude.

# 8. Auteur
Kadiatou Sadio Diallo



