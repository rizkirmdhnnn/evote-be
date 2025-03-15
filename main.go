package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/facades"

	"evote-be/bootstrap"
)

// @title evote-be API
// @version 1.0
// @description This is a sample server evote-be server
//
// @contact.name API Support
// @contact.email me@rizkirmdhn.cloud
// @host localhost:3000
// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @BasePath /
func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Start schedule by facades.Schedule
	go facades.Schedule().Run()

	// Start queue worker by facades.Queue()
	go func() {
		if err := facades.Queue().Worker().Run(); err != nil {
			facades.Log().Errorf("Queue run error: %v", err)
		}
	}()

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
		}
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}
		if err := facades.Schedule().Shutdown(); err != nil {
			facades.Log().Errorf("Schedule Shutdown error: %v", err)
		}

		os.Exit(0)
	}()

	select {}
}
