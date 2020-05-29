package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/werbot/lime/config"
	"github.com/werbot/lime/server/controllers"
	"github.com/werbot/lime/server/middleware"
)

// Start is a ...
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
	r.Use(sessions.Sessions(config.Config().GetString("cookie_name"), sessions.NewCookieStore([]byte(config.Config().GetString("cookie_secret")))))
	r.Static("/assets", "./server/web/assets")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate":  formatAsDate,
		"columnStatus":  columnStatus,
		"bytesToString": keyBytesToString,
	})
	r.LoadHTMLGlob("./server/web/templates/*.html")

	r.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	api := r.Group("/api")
	api.POST("/key", controllers.CreateKey)
	api.GET("/key/:customer_id", controllers.GetKey)
	api.PATCH("/key/:customer_id", controllers.UpdateKey)
	api.POST("/verify", controllers.VerifyKey)

	admin := r.Group("/admin")
	admin.GET("/", controllers.MainHandler)
	admin.POST("/login", middleware.Login)
	admin.POST("/logout", middleware.Logout)
	admin.Use(middleware.AuthRequired)
	{
		admin.GET("/subscription/:id/*action", controllers.CustomerSubscrptionList)
		admin.GET("/license/:id", controllers.DownloadLicense)
	}

	return r
}
