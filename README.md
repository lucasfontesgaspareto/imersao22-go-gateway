go install -tags 'postgres' github.comgolang-migrate/migrate/v4/cmd/migrate@latest

migrate -database "postgres://postgres:postgres@localhost:5432/gateway?sslmode=disable" -path migrations up