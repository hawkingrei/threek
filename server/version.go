package server

import (
	"github.com/hawkingrei/threek/internal/version"

	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": version.GitCommit,
	})
}
