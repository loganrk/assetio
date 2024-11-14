package gin

import (
	"assetio/internal/port"
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type routeGroup struct {
	ginGroup  *gin.RouterGroup
	accessLog port.Logger
}

// NewGroup creates a new route group with a specified groupName (path prefix)
func (r *route) NewGroup(groupName string) port.RouterGroup {
	return &routeGroup{
		ginGroup:  r.gin.Group(groupName),
		accessLog: r.accessLog,
	}
}

// RegisterRoute registers a new HTTP route with a specified method, path, and handler function.
func (r *routeGroup) RegisterRoute(method, path string, handlerFunc http.HandlerFunc) {
	// Register the route handler for the given method and path
	r.ginGroup.Handle(method, path, func(c *gin.Context) {

		// Wrap Gin's writer with our custom response writer to capture response details
		respWriter := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
			headers:        make(http.Header),
		}

		// Check if the initial response status is OK before proceeding
		if c.Writer.Status() == http.StatusOK {
			c.Writer = respWriter

			// Execute the handler function for the route
			handlerFunc(c.Writer, c.Request)

			// If the response is successful, log the success details
			if respWriter.statusCode == http.StatusOK {
				r.accessLog.Infow(c, "api response success",
					"method", c.Request.Method,
					"url", c.Request.URL.Path+"?"+c.Request.URL.RawQuery,
					"client-ip", c.ClientIP(),
					"headers", respWriter.headers,
				)
				return
			}
		}

		// Capture request body for logging, in case of failed responses
		var requestBody string
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore body for further use
			}
		}

		// Log details of failed API response (status, body, headers, etc.)
		r.accessLog.Warnw(c, "api response failed",
			"method", c.Request.Method,
			"url", c.Request.URL.Path+"?"+c.Request.URL.RawQuery,
			"request-body", requestBody,
			"status", respWriter.statusCode,
			"response", respWriter.body.String(),
			"headers", respWriter.headers,
			"client-ip", c.ClientIP(),
		)
	})
}

// UseBefore applies middleware to all routes within the route group.
func (r *routeGroup) UseBefore(middlewares ...http.Handler) {
	for _, middleware := range middlewares {
		r.ginGroup.Use(r.wrapHTTPHandlerFunc(middleware))
	}
}

// wrapHTTPHandlerFunc wraps the provided HTTP handler into a Gin handler function.
func (r *routeGroup) wrapHTTPHandlerFunc(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Execute the provided middleware handler
		h.ServeHTTP(c.Writer, c.Request)

		// If the response status is OK, proceed with the next handler
		if c.Writer.Status() == http.StatusOK {
			c.Next() // Continue to the next middleware or handler
		}
	}
}
