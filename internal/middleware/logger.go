package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrapperWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapper := &wrapperWriter{w, http.StatusOK}

		next.ServeHTTP(wrapper, r)

		slog.Info("Request completed", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start), "code", wrapper.statusCode)
	})
}
