SUITE – Module de statistiques des publications

1. Présentation

Ce projet a été réalisé dans le cadre du test technique SUITE.

Il permet de consulter les statistiques des publications d'un client sur quatre réseaux sociaux : Instagram, Facebook, TikTok et LinkedIn.

L'application affiche les impressions, les interactions, le taux d'engagement et les meilleures publications.

2. Technologies utilisées

Go : développement de l'API REST.

PostgreSQL : stockage des clients et des publications.

pgx : connexion entre Go et PostgreSQL.

Next.js et TypeScript : développement de l'interface.

Chart.js : affichage des statistiques sous forme de graphique.

Docker Compose : lancement des différents services.

golang-migrate : gestion des migrations SQL.

3. Lancement du projet sans dockerisation complète

Dans cette première approche, les composants sont lancés séparément.

Prérequis

Go installé.

Node.js et npm installés.

PostgreSQL installé et démarré.

Étape 1 : préparer la base de données

Créer un utilisateur et une base de données PostgreSQL :

CREATE USER suite_user WITH PASSWORD 'suite_password';
CREATE DATABASE suite_db OWNER suite_user;

Exécuter ensuite le script SQL initial depuis la racine du projet :

psql -U suite_user -d suite_db -f database/init.sql

Étape 2 : lancer l'API Go

Dans un terminal :

cd backend
go mod download
go test ./... -v

Configurer l'adresse de la base de données dans Git Bash :

export DATABASE_URL="postgres://suite_user:suite_password@localhost:5432/suite_db?sslmode=disable"

Démarrer l'API :

go run .

L'API est accessible sur http://localhost:8080.

Étape 3 : lancer le frontend

Dans un deuxième terminal :

cd frontend
npm install
npm run dev

L'application est accessible sur http://localhost:3000.

Exemple de tableau de bord :

http://localhost:3000/clients/1

4. Lancement avec Docker Compose

Pour simplifier le démarrage, j'ai ensuite dockerisé le projet.

Docker Compose permet de lancer PostgreSQL, l'API Go et le frontend avec une seule commande.

Démarrer les services

À la racine du projet :

docker compose up -d --build

Cette commande construit les images nécessaires et démarre les conteneurs.

Vérifier les services

docker compose ps -a

Vérifier les migrations

docker compose logs migrate

Les migrations sont exécutées automatiquement avant le démarrage de l'API.

Consulter les logs de l'API

docker compose logs api

Accéder à l'application

Frontend : http://localhost:3000/clients/1

API : http://localhost:8080

PostgreSQL : localhost:5439

Arrêter les services

docker compose down

Cette commande arrête les conteneurs sans supprimer les données de la base.

Le volume PostgreSQL est conservé. Il ne faut pas utiliser docker compose down -v si l'on souhaite garder les données.

5. Endpoints de l'API

Méthode

Endpoint

Description

GET

/clients/{id}/stats

Statistiques d'un client

GET

/clients/{id}/publications/top

Meilleures publications

Exemples :

curl "http://localhost:8080/clients/1/stats?from=2026-10-01&to=2026-10-31"

curl "http://localhost:8080/clients/1/publications/top?limit=3"

Le filtre par dates est facultatif. La limite des meilleures publications est de 5 par défaut et ne peut pas dépasser 20.

6. Tests

Les tests du backend sont exécutés avec :

cd backend
go test ./... -v

Ils vérifient notamment :

Le calcul du taux d'engagement.

Le cas où les impressions sont nulles.

Le fonctionnement du handler HTTP.

Le rejet d'un identifiant invalide.

J'ai également effectué des tests manuels des endpoints avec curl.

7. Choix techniques

J'ai choisi Go avec net/http pour construire une API simple, sans ajouter un framework dont je n'avais pas besoin.

PostgreSQL permet de stocker les clients et leurs publications. J'ai utilisé pgx pour communiquer avec la base de données.

Pour l'interface, j'ai utilisé Next.js et TypeScript. Chart.js permet de représenter les performances par plateforme.

Enfin, Docker Compose facilite le lancement de toute l'application. Les migrations SQL permettent de créer et de faire évoluer la structure de la base de données de manière organisée.

8. Ce que je n'ai pas eu le temps de réaliser

Je n'ai pas réalisé le bonus de génération automatique de rapports avec une intelligence artificielle.

Si j'avais eu plus de temps, j'aurais ajouté un endpoint Go qui récupère les statistiques du client.


9. Temps réellement consacré

Environ 4 jours ont été consacrés à ce test. Ce temps comprend le développement, les recherches et l'apprentissage des bases de Go à travers la documentation officielle, des tutoriels et des outils d'assistance comme Claude.

10. Auteur

Kadiatou Sadio Diallo