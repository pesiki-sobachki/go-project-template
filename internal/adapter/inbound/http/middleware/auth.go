package middleware

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shanth1/template/internal/pkg/response"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

// --- Basic Auth ---

func BasicAuth(user, pass string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()
			if !ok {
				basicAuthFailed(w)
				return
			}

			userMatch := subtle.ConstantTimeCompare([]byte(u), []byte(user)) == 1
			passMatch := subtle.ConstantTimeCompare([]byte(p), []byte(pass)) == 1

			if !userMatch || !passMatch {
				basicAuthFailed(w)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func basicAuthFailed(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	response.WithError(w, http.StatusUnauthorized, errors.New("unauthorized"))
}

// --- API Key Auth ---

func APIKeyAuth(validKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientKey := r.Header.Get("X-API-Key")

			if subtle.ConstantTimeCompare([]byte(clientKey), []byte(validKey)) != 1 {
				response.WithError(w, http.StatusUnauthorized, errors.New("invalid api key"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// --- JWT Auth ---

func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.WithError(w, http.StatusUnauthorized, errors.New("missing authorization header"))
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				response.WithError(w, http.StatusUnauthorized, errors.New("invalid authorization header format"))
				return
			}

			tokenString := parts[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				response.WithError(w, http.StatusUnauthorized, errors.New("invalid token"))
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if exp, ok := claims["exp"].(float64); ok {
					if time.Unix(int64(exp), 0).Before(time.Now()) {
						response.WithError(w, http.StatusUnauthorized, errors.New("token expired"))
						return
					}
				}

				ctx := r.Context()
				if sub, ok := claims["sub"].(string); ok {
					ctx = context.WithValue(ctx, UserIDKey, sub)
				}
				if role, ok := claims["role"].(string); ok {
					ctx = context.WithValue(ctx, RoleKey, role)
				}

				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				response.WithError(w, http.StatusUnauthorized, errors.New("invalid token claims"))
			}
		})
	}
}
