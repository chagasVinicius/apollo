package mid

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/chagasVinicius/apollo/kit/web"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Panics recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panics() web.Middleware {
	panicsCount := promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "web_panics_total",
		},
	)

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			// Defer a function to recover from a panic and set the err return
			// variable after the fact.
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					err = fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))
				}

				// inc prometheus metrics
				panicsCount.Inc()
			}()

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
