package middleware

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/misalud-ai/go-nurse/milog"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

const (
	masterOrgID = "misalud-organization"
)

type OAuth2 struct {
	authHost    string
	globalOrgID string
}

func NewOAuth2(authHost, globalOrgID string) *OAuth2 {
	return &OAuth2{
		authHost,
		globalOrgID,
	}
}

// https://auth0.com/docs/quickstart/backend/golang/01-authorization
func (m *OAuth2) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issuerURL, err := url.Parse(m.authHost)
		if err != nil {
			milog.ErrorWithError(r.Context(), err, "failed to parse the issuer url")
		}

		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			milog.Warnf(r.Context(), "authorization header is empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if accessToken != "" {
			if len(strings.Split(accessToken, " ")) < 2 {
				milog.Warnf(r.Context(), "authorization header doesn't have the right format")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			accessToken = strings.Split(accessToken, " ")[1]
		}

		if accessToken == "" {
			milog.Warnf(r.Context(), "authorization header is empty")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		pathParams := mux.Vars(r)
		orgID := pathParams["organization_id"]
		jwtParser := jwt.Parser{}
		token, _, _ := jwtParser.ParseUnverified(accessToken, jwt.MapClaims{})
		claims := token.Claims.(jwt.MapClaims)

		// if tokens issued with the global misalud or then authorize
		var audience string
		if claims["aud"] == masterOrgID {
			audience = masterOrgID
		} else if orgID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			audience = orgID
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{audience},
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			milog.WarnWithError(r.Context(), err, "failed to set up the jwt validator")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err = jwtValidator.ValidateToken(r.Context(), accessToken)
		if err != nil {
			milog.WarnWithError(r.Context(), err, "authorization header is invalid")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
