package jwt

import (
	"fmt"
	"strings"

	"github.com/f1rezy/book-store/internal/controllers"
	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		token = strings.Replace(token, "Bearer ", "", -1)
		fmt.Println(token)
		if token == "" {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		_, err := controllers.ParseToken(token)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
