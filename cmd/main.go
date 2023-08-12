package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil && shouldReadEnvFile() {
		e.Logger.Fatalf("Error loading .env file: %v", err)
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func shouldReadEnvFile() bool {
	return os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == ""
}
