package handler

import (
	"context"
	"fmt"
	"go-alodokter/config/env"
	"go-alodokter/utl/middleware/logging"
	"go-alodokter/utl/middleware/secure"
	"net/http"

	"os"
	"os/signal"
	"time"

	"go-alodokter/pkg/swagger"
	_ "go-alodokter/pkg/swagger/docs"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http

func (s *Service) HTTPServerMain() *echo.Echo {
	// to active swagger
	swagger.Init()

	e := echo.New()

	// e.Use(middleware.ServerHeader, middleware.Logger)

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.Use(secure.Headers())
	e.Use(logging.Logging())

	// administrator group
	adm := e.Group("/api/v1")

	// auth module
	ModuleAuth := adm.Group("/auth")
	s.AuthModule.HandleRest(ModuleAuth)

	// module identity access management
	ModuleUser := adm.Group("/user", s.MiddlewareAuth.BearerVerify())
	s.UserModule.HandleRest(ModuleUser)

	e.GET("/docs/*", echoSwagger.WrapHandler)

	return e
}

func (s *Service) StartServer() {
	server := s.HTTPServerMain()
	listenerPort := fmt.Sprintf(":%v", env.Conf.HttpPort)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func (s *Service) ShutdownServer() {
	server := s.HTTPServerMain()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
