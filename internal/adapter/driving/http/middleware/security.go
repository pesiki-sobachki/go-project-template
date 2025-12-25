package middleware

import (
	"errors"
	"net/http"

	"github.com/shanth1/template/internal/pkg/response"
)

func RequestSizeLimit(bytes int64) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ContentLength > bytes {
				response.WithError(w, http.StatusRequestEntityTooLarge, errors.New("request body too large"))
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, bytes)
			next.ServeHTTP(w, r)
		})
	}
}
