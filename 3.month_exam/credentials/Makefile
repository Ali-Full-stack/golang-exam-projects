
	migrate_import:
		go get github.com/golang-migrate/migrate/v4/database/postgres

	migrate_create:
		migrate create -ext sql -dir migrations -seq credential

	migrate_up:
		migrate -database postgres://postgres:3115@localhost:5432/credentials?sslmode=disable -path ./migrations up

	migrate_down:
		migrate -database postgres://postgres:3115@localhost:5432/credentials?sslmode=disable -path ./migrations migrate_down

	migrate_force:
		migrate -database postgres://postgres:3115@localhost:5432/credentials?sslmode=disable -path ./migrations force

	client_proto:
		protoc --go_out=. --go-grpc_out=. protos/clientpb/client.proto


