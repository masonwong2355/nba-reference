APP_ENV ?= local
DATETIME=$(shell date -u "+%Y%m%d%H%M%S")

dep:
	go mod tidy

run:
	APP_ENV=$(APP_ENV) go run cmd/api/main.go

init-db:
	@echo "Initializing DB..."
	# Replace with your actual DB init command, e.g., docker-compose up -d db

migrate:
	@echo "Migrating DB..."
	# Add your migration tool here, e.g. goose, migrate, etc.

build:
	go build -o nba-api cmd/api/main.go

deploy:
	@echo "Deploying app..."
	# Add deploy commands (scp, docker, etc.)

lint:
	golangci-lint run ./...



DB_URL="postgres://postgres:password@localhost:5432/nba_dev?sslmode=disable"

# migrate -database "postgres://postgres:password@localhost:5432/nba_dev?sslmode=disable" -path ./database/migrations force 20250725120718

migrate-up:
	migrate -database "$(DB_URL)" -path ./database/migrations/ up

migrate-down:
	migrate -database "$(DB_URL)" -path ./database/migrations down 1

gen-migration-file:
	printf "Enter file names: "; read -r FILE_NAME; \
	migrate create -ext sql -dir database/migrations $$FILE_NAME