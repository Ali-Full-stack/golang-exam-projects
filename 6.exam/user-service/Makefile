
migrate_import:
	go get github.com/golang-migrate/migrate/v4/database/postgres

migrate_create:
	migrate create -ext sql -dir migrations -seq users

migrate_up:
	migrate -database postgres://postgres:7777@localhost:5432/users?sslmode=disable -path ./migrations up

migrate_down:
	migrate -database postgres://postgres:7777@localhost:5432/users?sslmode=disable -path ./migrations down

migrate_force:
	migrate -database postgres://postgres:7777@localhost:5432/users?sslmode=disable -path ./migrations force

sqlc-generate:
	sqlc vet ; sqlc generate

gen-protoc:
	protoc --go_out=. --go-grpc_out=. ./proto/user.proto
