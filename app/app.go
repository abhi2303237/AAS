package app

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Config struct {
	EnableMetricsUi bool
	AppTitle        string
	ContextPath     string
	LogLevel        string
}

func NewFiberApp(conf fiber.Config, applicationConf Config) *fiber.App {
	fiber := fiber.New(conf)
	setDefaultMiddlewares(fiber, applicationConf.AppTitle, applicationConf.EnableMetricsUi)

	switch strings.ToLower(applicationConf.LogLevel) {
	case "debug":
		log.SetLevel(log.LevelDebug)
	case "info":
		log.SetLevel(log.LevelInfo)
	case "warn":
		log.SetLevel(log.LevelWarn)
	case "error":
		log.SetLevel(log.LevelError)
	case "fatal":
		log.SetLevel(log.LevelFatal)
	case "panic":
		log.SetLevel(log.LevelPanic)
	default:
		log.SetLevel(log.LevelError) // Default to Error if logLevel is not recognized
	}

	return fiber
}

func setDefaultMiddlewares(fiberApp *fiber.App, appTitle string, enableMetricsUi bool) {
	fiberApp.Get("/metrics", monitor.New(
		monitor.Config{
			Title:   appTitle,
			Refresh: 1,
			APIOnly: !enableMetricsUi,
		},
	))
	fiberApp.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
	}))
	fiberApp.Use(recover.New())
	fiberApp.Use(requestid.New(requestid.Config{
		Header: "X-IApps-Request-Id",
		// Generator: func() string {
		// 	return "static-id"
		// },
	}))
	fiberApp.Use(logger.New(logger.Config{
		Format:     "${time} ${locals:requestid} ${status} - ${method} ${path}\n",
		TimeFormat: time.DateTime,
		TimeZone:   "Local",
	}))

}
