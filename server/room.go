package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/threek/store"
)

func GetRoom(c *gin.Context) {
	//username, _ := c.Get("user")
	rid, _ := c.Get("room")
	user, err := store.GetRoom(c, rid.(int))
	if err != nil {
		c.String(500, "Error fetching user. %s", err)
		return
	}
	c.JSON(200, user)
}
