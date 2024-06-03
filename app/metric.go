package app

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

var apiSummary = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "http_server_requests_seconds",
		Help:       "Duration of HTTP server request handling",
		Objectives: map[float64]float64{0.5: 0.05, 0.95: 0.005},
	},
	[]string{"method", "status", "path"},
)

func init() {
	prometheus.MustRegister(apiSummary)
}

func MetricHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func MetricMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		apiSummary.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(status),
			c.Request.URL.Path,
		).Observe(latency.Seconds())
	}
}
