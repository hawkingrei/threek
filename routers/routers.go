package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hawkingrei/threek/server"
)

func Load(middleware ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	//authMiddleware := &auth.GinAuthMiddleware{}
	e.Use(gin.Recovery())
	e.Use(middleware...)
	version := e.Group("/api/version")
	{
		version.GET("", server.GetVersion)
	}
	room := e.Group("/api/room")
	{
		//room.POST("", server.CreateRoom)
		room.GET("", server.GetRoom)
		//room.GET("/cancel", server.CancelRoom)
	}
	return e
}
