DB_URL=postgresql://postgres:postgres@localhost:5432/test_db?sslmode=disable

make_migration:
	migrate create -ext sql -dir migrations -seq $(name)

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

startpg:
	docker start postgres