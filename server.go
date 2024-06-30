package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thanapongsj1996/assessment/modules/expense/handler"
	"github.com/thanapongsj1996/assessment/modules/expense/repository"
	"github.com/thanapongsj1996/assessment/modules/expense/service"
	"github.com/thanapongsj1996/assessment/routes"

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
	db := database.GetDB()
	defer database.CloseDB()

	expenseRepo := repository.NewExpenseRepository(db)
	expenseService := service.NewExpenseService(expenseRepo)
	expenseHandler := handler.NewExpenseHandler(expenseService)

	// Init routes
	routes.NewExpenseRoute(e, expenseHandler)

	// health
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
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
