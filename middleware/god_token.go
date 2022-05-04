package middleware

import (
	"net/http"
	"strings"
)

type GodToken struct {
	godTokens []string
}

func NewGodToken(godTokens []string) *GodToken {
	return &GodToken{
		godTokens,
	}
}

func (m *GodToken) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		authTokens := r.Header.Values(HeaderXMSToken)

		if len(authTokens) > 1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authToken := strings.TrimSpace(authTokens[0])

		for _, godToken := range m.godTokens {
			if authToken == godToken {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.WriteHeader(http.StatusUnauthorized)
	})
}
