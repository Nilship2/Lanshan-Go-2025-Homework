package middleware

import (
	"lesson7/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
