package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	augLogger = log.New(&logWriter{os.Stdout}, " | [Augmentor] | ", 0)
)

// loggingHandler is an http middleware that logs information about its children.
// It logs the status code, latency, method, path, and the reason for returning a
// 204 if provided.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)
		augLogger.Printf("%3d | %13v | %-7s %s | %s",
			lrw.statusCode,
			time.Since(start),
			r.Method,
			r.URL.Path,
			lrw.Header().Get("X-Pass-Reason"),
		)
	})
}

// loggingResponseWriter stores the status code being written out.
// This is useful for loggingHandler, which retrieves it later.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader stores the status code before calling WriteHeader
// on the http.ResponseWriter provided initially.
func (w *loggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// logWriter is an io.Writer that logs the time, in RFC3339 format,
// when Write is called.
type logWriter struct {
	out io.Writer
}

// Write writes the time in RFC3339 format to the underlying
// io.Writer before writing the passed bytes.
func (w *logWriter) Write(p []byte) (n int, err error) {
	t := []byte(time.Now().Format(time.RFC3339))
	n, err = w.out.Write(t)
	if err != nil {
		return
	}
	n, err = w.out.Write(p)
	return
}
