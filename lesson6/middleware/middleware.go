// middleware/middleware.go
package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Example1() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		log.Printf("中间件1 | %3d | %13v | %15s | %s | %s",
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			c.Request.Method,
			path,
		)
	}
}
func Example2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ping-ping-pong-pong！
		time.Sleep(114 * time.Millisecond)
		c.Next()
		log.Printf("乒乒乓乓完成！")
	}
}
