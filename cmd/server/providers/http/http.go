package http

import (
	"fmt"
	"watermap/cmd/server/providers"
	"watermap/infrastructure/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type httpServer struct {
	repository db.Repository
}

func (*httpServer) Serve() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("localhost:%d", viper.GetInt("http_port")))
}
func Provider() providers.Server {
	return &httpServer{}
}
