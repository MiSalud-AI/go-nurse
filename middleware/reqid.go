package middleware

import (
	"context"
	"net/http"

	"github.com/misalud-ai/go-nurse/milog"

	"github.com/google/uuid"
)

type RequestID struct{}

func NewRequestID() *RequestID {
	return &RequestID{}
}

func (m *RequestID) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(HeaderNameRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), milog.ContextKeyRequestID, requestID)
		w.Header().Add(HeaderNameRequestID, requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
