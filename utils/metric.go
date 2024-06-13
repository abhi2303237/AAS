package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func InitMetrics(router fiber.Router) {
	conf := GetConfig()
	router.Get("/metrics", monitor.New(
		monitor.Config{
			Title:   "DataStream Metrics",
			Refresh: 1,
			APIOnly: !conf.EnableMetricsUi,
		},
	))
}
