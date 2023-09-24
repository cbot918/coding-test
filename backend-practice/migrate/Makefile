DB_URL=postgresql://postgres:12345@localhost:5433/practice?sslmode=disable
MIG_PATH=db/migrations

.PHONY: postgres
newdb:
	docker run --name practice-psql -p 5433:5432 -e POSTGRES_DB=practice -e POSTGRES_PASSWORD=12345 -d postgres

.PHONY: mig-init
migadd:
	migrate create -ext sql -dir db/migrations -seq $(name)

migup:
	migrate -database $(DB_URL) -path "$(MIG_PATH)" up

migup1:
	migrate -database $(DB_URL) -path "$(MIG_PATH)" up 1

migdown:
	migrate -database $(DB_URL) -path "$(MIG_PATH)" down

migdown1:
	migrate -database $(DB_URL) -path "$(MIG_PATH)" down 1