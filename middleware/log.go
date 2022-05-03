package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/misalud-ai/go-nurse/milog"
	log "github.com/sirupsen/logrus"
)

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

		lw := newResponseWriterWrapper(w)

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
