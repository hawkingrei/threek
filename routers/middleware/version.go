package middleware

import (
	"github.com/hawkingrei/threek/internal/version"

	"github.com/gin-gonic/gin"
)

// Version is a middleware function that appends the Redp
// version information to the HTTP response. This is intended
// for debugging and troubleshooting.
func Version(c *gin.Context) {
	c.Header("X-REDP-VERSION", version.GitCommit)
	c.Next()
}
