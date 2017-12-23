package server

import (
	"strconv"

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

func CreateRoom(c *gin.Context) {
	room, err := store.CreateRoom(c)
	if err != nil {
		c.String(500, "Error fetching user. %s", err)
		return
	}
	c.JSON(200, room)
}

func JoinRoom(c *gin.Context) {
	strRid := c.Request.Header.Get("roomid")
	username, _ := c.Get("user")
	rid, err := strconv.ParseInt(strRid, 10, 64)
	if err != nil {
		c.String(500, "rid error. %s", err)
		return
	}
	rm, err := store.JoinRoom(c, rid, username.(string))
	if err != nil {
		c.String(500, "Error fetching user. %s", err)
		return
	}
	c.JSON(200, rm)
}
