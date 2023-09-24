
# databases

MIG_DB_URL=postgres://postgres:12345@localhost:5433/practice2?sslmode=disable
MIG_PATH=db/migrations
db:
	docker run 

miginit:
	migrate create -ext sql -dir db/migrations -seq create_users_table

migadd:
	migrate create -ext sql -dir db/migrations -seq $(name)

migup:
	migrate -database $(MIG_DB_URL) -path $(MIG_PATH) up
migup1:
	migrate -database $(MIG_DB_URL) -path $(MIG_PATH) up 1

migdown:
	migrate -database $(MIG_DB_URL) -path $(MIG_PATH) down
migdown1:
	migrate -database $(MIG_DB_URL) -path $(MIG_PATH) down 1


