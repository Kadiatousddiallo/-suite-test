# Étape 1 : compilation de l'application Go
FROM golang:1.27.3-alpine AS builder

WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le code source
COPY . .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o suite-api .

# Étape 2 : image légère pour exécuter l'API
FROM alpine:3.22

WORKDIR /app

# Ajouter les certificats nécessaires
RUN apk add --no-cache ca-certificates

# Copier le programme compilé
COPY --from=builder /app/suite-api .

# Port de l'API
EXPOSE 8080

# Lancer l'API
CMD ["./suite-api"]