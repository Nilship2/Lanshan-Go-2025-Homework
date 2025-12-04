// middleware/middleware.go
package middleware

import (
	"log"
	"net/http"
	"nilmod/utils"
	"strings"
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
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		username, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}
		// 在 context 中记录用户名，后续 handler 可取出
		c.Set("username", username)
		//fmt.Println("<<<", username)
		c.Next()
	}
}
