package gin

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body       *bytes.Buffer // Buffer to store response body data
	statusCode int           // Store the HTTP status code
	headers    http.Header   // Store the response headers
}

// Write intercepts the response body data and writes it to both the buffer and the client
func (w *responseWriter) Write(data []byte) (int, error) {
	w.body.Write(data)                  // Write to buffer for logging
	return w.ResponseWriter.Write(data) // Write to actual ResponseWriter
}

// WriteHeader intercepts the status code and headers
func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode                // Store the status code
	w.headers = w.Header().Clone()           // Clone headers for logging
	w.ResponseWriter.WriteHeader(statusCode) // Write headers to actual ResponseWriter
}
