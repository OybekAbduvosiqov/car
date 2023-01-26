migration-up:
	migrate -path ./migration/postgres/ -database 'postgres://oybek:oybek@localhost:5432/exam?sslmode=disable' up 


migration-down:
	migrate -path ./migration/postgres/ -database 'postgres://oybek:oybek@localhost:5432/exam?sslmode=disable' down



run:
	go run cmd/main.go
swag:
	swag init -g api/api.go -o api/docs

