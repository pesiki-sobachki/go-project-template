package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shanth1/gotools/log"
	"github.com/shanth1/gotools/logkeys"
)

func Logger(l log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			reqID := middleware.GetReqID(r.Context())
			ctxLogger := l.With(log.Str(logkeys.RequestID, reqID))
			ctx := log.NewContext(r.Context(), ctxLogger)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r.WithContext(ctx))

			l.Info().
				Str(logkeys.RequestID, reqID).
				Str(logkeys.HTTPMethod, r.Method).
				Str(logkeys.HTTPPath, r.URL.Path).
				Int(logkeys.HTTPStatus, ww.Status()).
				Int(logkeys.BytesOut, ww.BytesWritten()).
				Dur(logkeys.Latency, time.Since(start)).
				Str(logkeys.RemoteAddr, r.RemoteAddr).
				Msg("http_request")
		})
	}
}
