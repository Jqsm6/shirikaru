migrate_force:
	migrate -database postgres://postgres:postgres@localhost:5432/anime_db?sslmode=disable -path migrations force 1

migrate_version:
	migrate -database postgres://postgres:postgres@localhost:5432/anime_db?sslmode=disable -path migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@localhost:5432/anime_db?sslmode=disable -path migrations up 1

migrate_down:
	migrate -database postgres://postgres:postgres@localhost:5432/anime_db?sslmode=disable -path migrations down 1

run-linter:
	golangci-lint run

run:
	go run ./cmd/api/main.go

build:
	go build -o ./bin/main ./cmd/api/main.go