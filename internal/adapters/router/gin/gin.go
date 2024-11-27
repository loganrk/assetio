package gin

import (
	"assetio/internal/port"
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	gin       *gin.Engine
	accessLog port.Logger
}

// New creates a new instance of the Gin router with a custom access logger.
func New(accessLoggerIns port.Logger) port.Router {
	gin.DisableConsoleColor()

	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                             // Allow specific origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  // Allowed methods
		AllowHeaders:     []string{"Authorization", "Content-Type"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                // Headers exposed to the browser
		AllowCredentials: true,                                      // Allow credentials (cookies, Authorization header)
		MaxAge:           12 * time.Hour,                            // Preflight cache duration
	}))
	return &route{
		gin:       r,
		accessLog: accessLoggerIns,
	}
}

// RegisterRoute registers a new HTTP route with a specified method, path, and handler function.
func (r *route) RegisterRoute(method, path string, handlerFunc http.HandlerFunc) {
	r.gin.Handle(method, path, func(c *gin.Context) {
		// Create a custom response writer to capture response status, body, and headers
		respWriter := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
			headers:        make(http.Header),
		}
		c.Writer = respWriter

		// Execute the handler function
		handlerFunc(c.Writer, c.Request)

		// Log the response if status code is OK (200)
		if respWriter.statusCode == http.StatusOK {
			r.accessLog.Infow(c, "api response success",
				"method", c.Request.Method,
				"url", c.Request.URL.Path+"?"+c.Request.URL.RawQuery,
				"client-ip", c.ClientIP(),
				"headers", respWriter.headers,
			)
		} else {
			// Log failed response with additional details
			var requestBody string
			// Capture request body if POST, PUT, or PATCH request
			if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
				bodyBytes, err := io.ReadAll(c.Request.Body)
				if err == nil {
					requestBody = string(bodyBytes)
					c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}
			}

			// Log the failure with response body, status code, and request body
			r.accessLog.Warnw(c, "api response failed",
				"method", c.Request.Method,
				"url", c.Request.URL.Path+"?"+c.Request.URL.RawQuery,
				"request-body", requestBody,
				"status", respWriter.statusCode,
				"response", respWriter.body.String(),
				"headers", respWriter.headers,
				"client-ip", c.ClientIP(),
			)
		}
	})
}

// StartServer starts the Gin HTTP server on the specified port.
func (r *route) StartServer(port string) error {
	// Run the Gin server on the specified port
	return r.gin.Run(":" + port)
}

// UseBefore registers middleware to be executed before the route handler.
func (r *route) UseBefore(middlewares ...http.Handler) {
	for _, middleware := range middlewares {
		// Wrap each middleware and add it to the Gin router
		r.gin.Use(r.wrapHTTPHandlerFunc(middleware))
	}
}

// wrapHTTPHandlerFunc wraps an HTTP handler to work with Gin's context.
func (r *route) wrapHTTPHandlerFunc(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Execute the middleware (http.Handler)
		h.ServeHTTP(c.Writer, c.Request)

		// Proceed to the next handler if the status is OK
		if c.Writer.Status() == http.StatusOK {
			c.Next() // Continue to the next middleware or handler
		}
	}
}
