package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Cache() func(c *gin.Context) {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if uri == "/" || strings.HasPrefix(uri, "/api") || strings.HasPrefix(uri, "/v1") {
			c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
		} else {
			c.Header("Cache-Control", "max-age=604800") // one week
		}
		c.Next()
	}
}
