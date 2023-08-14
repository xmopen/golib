// Package middleware provide middlewares.
package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors Cross Origin Resource Sharing is an HTTP-header based mechanism that allows a server to indicate any origins
// other than its own from a browser should permit loading resources.
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:    []string{"Content-Type"},
		MaxAge:          time.Hour * 24 * 7, // cros 缓存7天,前提浏览器开启缓存.
	})
}
