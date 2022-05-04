package middleware

import (
	"net/http"
	"strings"
)

type GodTokenOnly struct {
	godTokens []string
}

func NewGodTokenOnly(godTokens []string) *GodTokenOnly {
	return &GodTokenOnly{
		godTokens,
	}
}

func (m *GodTokenOnly) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if len(m.godTokens) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		authTokens := r.Header.Values(HeaderXMSToken)

		if len(authTokens) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

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
