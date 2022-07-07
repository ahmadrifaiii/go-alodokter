package swagger

import "go-alodokter/pkg/swagger/docs"

// @termsOfService http://swagger.io/terms/

func Init() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Go Lemonilo User Management"
	docs.SwaggerInfo.Description = "LEMONILO"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
