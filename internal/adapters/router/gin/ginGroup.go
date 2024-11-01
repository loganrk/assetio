package gin

import (
	"assetio/internal/port"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type routeGroup struct {
	ginGroup  *gin.RouterGroup
	accessLog port.Logger
}

func (r *route) NewGroup(groupName string) port.RouterGroup {
	return &routeGroup{
		ginGroup:  r.gin.Group(groupName),
		accessLog: r.accessLog,
	}
}

func (r *routeGroup) RegisterRoute(method, path string, handlerFunc http.HandlerFunc) {
	r.ginGroup.Handle(method, path, func(c *gin.Context) {

		// Wrap Gin's writer with our custom response writer
		respWriter := &responseWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
			headers:        make(http.Header), // Initialize the header storage
		}
		c.Writer = respWriter

		// Execute the handler
		handlerFunc(c.Writer, c.Request)

		fmt.Println(respWriter.statusCode)
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

func (r *routeGroup) UseBefore(middlewares ...http.Handler) {
	for _, middleware := range middlewares {
		r.ginGroup.Use(r.wrapHTTPHandlerFunc(middleware))
	}
}

func (r *routeGroup) wrapHTTPHandlerFunc(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Wrap the handler with the provided middleware
		h.ServeHTTP(c.Writer, c.Request)

	}
}
