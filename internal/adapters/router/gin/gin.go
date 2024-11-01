package gin

import (
	"assetio/internal/port"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type route struct {
	gin       *gin.Engine
	accessLog port.Logger
}

func New(accessLoggerIns port.Logger) port.Router {
	gin.DisableConsoleColor()
	return &route{
		gin:       gin.Default(),
		accessLog: accessLoggerIns,
	}
}

func (r *route) RegisterRoute(method, path string, handlerFunc http.HandlerFunc) {
	r.gin.Handle(method, path, func(c *gin.Context) {
		// Wrap Gin's writer with our custom response writer
		respWriter := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
			headers:        make(http.Header), // Initialize the header storage
		}
		c.Writer = respWriter

		// Execute the handler
		handlerFunc(c.Writer, c.Request)

		// Determine logging based on response status
		if respWriter.statusCode == http.StatusOK {
			r.accessLog.Infow(c, "api response success",
				"method", c.Request.Method,
				"url", c.Request.URL.Path+"?"+c.Request.URL.RawQuery,
				"client-ip", c.ClientIP(),
				"headers", respWriter.headers, // Log response headers
			)
		} else {
			// Capture request body if needed
			var requestBody string
			if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
				bodyBytes, err := io.ReadAll(c.Request.Body)
				if err == nil {
					requestBody = string(bodyBytes)
					c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore body for further use
				}
			}

			r.accessLog.Warnw(c, "api response failed",
				"method", c.Request.Method,
				"url", c.Request.URL.Path+"?"+c.Request.URL.RawQuery,
				"request-body", requestBody,
				"status", respWriter.statusCode, // Log status code
				"response", respWriter.body.String(), // Log response body
				"headers", respWriter.headers, // Log response headers
				"client-ip", c.ClientIP(),
			)
		}
	})
}

func (r *route) StartServer(port string) error {
	return r.gin.Run(":" + port)
}

func (r *route) UseBefore(middlewares ...http.Handler) {
	for _, middleware := range middlewares {
		r.gin.Use(r.wrapHTTPHandlerFunc(middleware))
	}
}

func (r *route) wrapHTTPHandlerFunc(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the handler with the provided middleware
		h.ServeHTTP(c.Writer, c.Request)

		if c.Writer.Status() == http.StatusOK {
			c.Next()
		}

	}
}
