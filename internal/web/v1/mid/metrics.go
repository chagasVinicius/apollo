package mid

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/chagasVinicius/apollo/kit/web"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics updates program counters.
func Metrics(buckets ...float64) web.Middleware {
	defaultBuckets := []float64{100, 300, 800, 1300, 5000}

	const (
		reqsName    = "web_requests_total"
		latencyName = "web_request_duration_milliseconds"
	)

	if len(buckets) == 0 {
		buckets = defaultBuckets
	}

	reqsCounter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: reqsName,
			Help: "How many HTTP requests processed, partitioned by status code, method and HTTP path (with patterns).",
		},
		[]string{"status_code", "status_class", "method", "path"},
	)

	latencyHist := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    latencyName,
			Help:    "How long it took to process the request, partitioned by status code, method and HTTP path.",
			Buckets: buckets,
		},
		[]string{"status_code", "status_class", "method", "path"},
	)

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			start := time.Now()

			err := handler(ctx, w, r)

			v := web.GetValues(ctx)
			routePattern := web.ContextRoute(ctx)
			statusCode := strconv.Itoa(v.StatusCode)

			var statusClass string
			switch {
			case v.StatusCode >= 500:
				statusClass = "5xx"
			case v.StatusCode >= 400:
				statusClass = "4xx"
			case v.StatusCode >= 300:
				statusClass = "3xx"
			default:
				statusClass = "2xx"
			}

			reqsCounter.WithLabelValues(
				statusCode,
				statusClass,
				r.Method,
				routePattern,
			)

			latencyHist.WithLabelValues(
				statusCode,
				statusClass,
				r.Method,
				routePattern,
			).Observe(float64(time.Since(start).Nanoseconds()) / 1000000)

			return err
		}

		return h
	}

	return m
}
