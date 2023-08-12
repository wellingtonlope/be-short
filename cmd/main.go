package main

import (
	"fmt"
	"github.com/wellingtonlope/be-short/internal/infra/mongo"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"github.com/wellingtonlope/be-short/pkg/log"
)

const applicationName = "be-short"

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil && isDevEnvironment() {
		e.Logger.Fatalf("Error loading .env file: %v", err)
	}
	enableNewRelic(e)

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	database, err := mongo.NewDatabase(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DATABASE"))
	if err != nil {
		e.Logger.Fatalf("Error loading database: %v", err)
	}

	_, err = mongo.NewShorted(database)
	if err != nil {
		e.Logger.Fatalf("Error loading database: %v", err)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func enableNewRelic(e *echo.Echo) {
	if isDevEnvironment() {
		return
	}
	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(fmt.Sprintf("%s.%s", os.Getenv("APP_ENV"), applicationName)),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		e.Logger.Errorf("Error loading newrelic: %v", err)
	}
	e.Use(nrecho.Middleware(newRelicApp))

	logLogrus := logrus.New()
	logLogrus.SetFormatter(nrlogrus.NewFormatter(newRelicApp, &logrus.TextFormatter{}))
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logLogrus.SetLevel(level)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetRequest(c.Request().WithContext(log.WithLogger(c.Request().Context(), logLogrus)))
			return next(c)
		}
	})
}

func isDevEnvironment() bool {
	return os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == ""
}
