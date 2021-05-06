package noFloc

import "github.com/gin-gonic/gin"

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// the special header to tell google to leave us alone
		c.Writer.Header().Set("Permissions-Policy", "interest-cohort=()")

		c.Next()
	}
}
