package middleware

import (
	"fmt"
	"net/http"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/gorilla/mux"
)

type Metrics struct {
	ddClient *statsd.Client
	clock    timer
}

func NewMetrics(ddClient *statsd.Client) *Metrics {
	return &Metrics{
		ddClient: ddClient,
		clock:    &realClock{},
	}
}

// Middleware implement mux middleware interface
func (m *Metrics) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := m.clock.Now()

		lw := newResponseWriterWrapper(w)

		next.ServeHTTP(lw, r)
		route := mux.CurrentRoute(r)

		latency := m.clock.Since(start)
		// service.app-name.request
		// service.app-name.response_time
		tags := []string{
			fmt.Sprintf("route:%s", route.GetName()),
			fmt.Sprintf("method:%s", r.Method),
			fmt.Sprintf("status_code:%d", lw.statusCode),
		}
		m.ddClient.Timing("response_time", latency, tags, 1)
		m.ddClient.Count("request", 1, tags, 1)
	})
}
