package main

import (
	"go-alodokter/config/env"
	"go-alodokter/handler"
	"os"
	"os/signal"
)

func main() {
	env.LoadEnv()

	// Init dependencies
	service := handler.InitHandler()

	// start echo server
	service.StartServer()

	// Shutdown with gracefull handler
	service.ShutdownServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
