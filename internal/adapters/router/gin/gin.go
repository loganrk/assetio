package gin

import (
	"assetio/internal/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type route struct {
	gin *gin.Engine
}

type routeGroup struct {
	ginGroup *gin.RouterGroup
}

func New() port.Router {
	gin.DisableConsoleColor()
	return &route{
		gin: gin.Default(),
	}
}

func (r *route) NewGroup(groupName string) port.RouterGroup {
	return &routeGroup{
		ginGroup: r.gin.Group(groupName),
	}
}

func (r *route) RegisterRoute(method, path string, handlerFunc http.HandlerFunc) {
	r.gin.Handle(method, path, func(c *gin.Context) {
		handlerFunc(c.Writer, c.Request)
	})
}

func (r *route) StartServer(port string) error {
	return r.gin.Run(":" + port)
}

func (r *route) UseBefore(middlewares ...http.HandlerFunc) {
	for _, middleware := range middlewares {
		r.gin.Use(wrapHTTPHandlerFunc(middleware))
	}
}

func (r *routeGroup) RegisterRoute(method, path string, handlerFunc http.HandlerFunc) {
	r.ginGroup.Handle(method, path, func(c *gin.Context) {
		handlerFunc(c.Writer, c.Request)
	})
}

func (r *routeGroup) UseBefore(middlewares ...http.HandlerFunc) {
	for _, middleware := range middlewares {
		r.ginGroup.Use(wrapHTTPHandlerFunc(middleware))
	}
}

func wrapHTTPHandlerFunc(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
		if c.Writer.Status() != http.StatusOK {
			c.Abort()
		}
		c.Next()
	}
}
