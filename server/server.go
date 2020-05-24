package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/werbot/lime/config"
	"github.com/werbot/lime/server/controllers"
)

func Start() {
	cfg := config.Config()
	gin.SetMode(cfg.GetString("mode"))
	r := setupRouter()

	//seed.Load(config.DB)

	err := r.Run(cfg.GetString("port"))
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/key", controllers.CreateKey)
	r.GET("/key/:customer_id", controllers.GetKey)
	r.PATCH("/key/:customer_id", controllers.UpdateKey)

	r.POST("/verify", controllers.VerifyKey)

	return r
}
