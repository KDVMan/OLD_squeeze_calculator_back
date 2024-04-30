package core_middlewares_logger

import (
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"time"
)

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := log.With(
			slog.String("component", "core/middlewares/logger"),
		)

		log.Info("logger middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)

			rw := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			timeStart := time.Now()

			defer func() {
				entry.Info("request completed",
					slog.Int("status", rw.Status()),
					slog.Int("bytes", rw.BytesWritten()),
					slog.String("duration", time.Since(timeStart).String()),
				)
			}()

			next.ServeHTTP(rw, r)
		}

		return http.HandlerFunc(fn)
	}
}
