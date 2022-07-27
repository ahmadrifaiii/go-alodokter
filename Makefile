init:
	@go mod init
	
tidy:
	@go mod tidy

swagger-gen:
	@swag init -g main.go --output pkg/swagger/docs

run-migrate:
	@go run utl/database/migrate/mysql/main.go