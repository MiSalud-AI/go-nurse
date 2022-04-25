package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/misalud-ai/go-nurse/milog"
	log "github.com/sirupsen/logrus"
)

type timer interface {
	Now() time.Time
	Since(time.Time) time.Duration
}

// realClock save request times
type realClock struct{}

func (rc *realClock) Now() time.Time {
	return time.Now()
}

func (rc *realClock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

type Logging struct {
	// logger *logrus.Logger
	clock timer
}

func NewLogger() *Logging {
	return &Logging{
		clock: &realClock{},
	}
}

// realIP get the real IP from http request
func realIP(req *http.Request) string {
	ra := req.RemoteAddr
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		ra = strings.Split(ip, ", ")[0]
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		ra = ip
	} else {
		ra, _, _ = net.SplitHostPort(ra)
	}
	return ra
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}

func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
	return lw.ResponseWriter.Write(b)
}

// Middleware implement mux middleware interface
func (m *Logging) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		entry := log.NewEntry(log.StandardLogger())
		start := m.clock.Now()

		if reqID := r.Header.Get(milog.HTTPHeaderNameRequestID); reqID != "" {
			entry = entry.WithField("requestId", reqID)
		}

		if remoteAddr := realIP(r); remoteAddr != "" {
			entry = entry.WithField("remoteAddr", remoteAddr)
		}

		entry = entry.WithFields(log.Fields{
			"request": r.RequestURI,
			"method":  r.Method,
		})

		lw := newLoggingResponseWriter(w)

		next.ServeHTTP(lw, r)

		ctx := r.Context()
		reqID := ctx.Value(milog.ContextKeyRequestID)

		if v, ok := reqID.(string); ok {
			entry = entry.WithField("requestId", v)
		}

		latency := m.clock.Since(start)

		entry.WithFields(log.Fields{
			"status": lw.statusCode,
			"took":   latency,
		}).Info("completed handling request")
	})
}
