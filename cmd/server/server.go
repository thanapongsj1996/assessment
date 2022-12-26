package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/subosito/gotenv"
	"github.com/thanapongsj1996/assessment/config"
	"github.com/thanapongsj1996/assessment/database"
)

func init() {
	gotenv.Load()
}

func main() {
	e := echo.New()

	// Environment config
	appConfig := config.AppConfig()

	// Init Database
	database.InitDB(appConfig.DatabaseUrl)
	//db := database.GetDB()
	defer database.CloseDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	go func() {
		if err := e.Start(":" + appConfig.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
