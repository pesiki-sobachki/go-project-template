package middleware

import (
	"net"
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

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			host, port, _ := net.SplitHostPort(r.RemoteAddr)
			l.Info().
				Str(logkeys.HTTPMethod, r.Method).
				Str(logkeys.HTTPPath, r.URL.Path).
				Int(logkeys.HTTPStatus, ww.Status()).
				Int(logkeys.BytesIn, ww.BytesWritten()).
				Dur(logkeys.Latency, time.Since(start)).
				Str(logkeys.RemoteIP, host).
				Str(logkeys.RemotePort, port).
				Msg("http_request")
		})
	}
}
