package middleware

import (
    "log"
    "net/http"
    "time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Wrap ResponseWriter to capture status code
        wrapped := wrapResponseWriter(w)
        next.ServeHTTP(wrapped, r)

        log.Printf(
            "%s %s %s %d %s",
            r.RemoteAddr,
            r.Method,
            r.URL.Path,
            wrapped.status,
            time.Since(start),
        )
    })
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
    return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}
