init:
	@go mod init
	
tidy:
	@go mod tidy

swagger-gen:
	@swag init -g main.go --output pkg/swagger/docs