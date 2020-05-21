package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("static", "./static")
	//r.LoadHTMLGlob("views/*")
	r.LoadHTMLFiles("views/login.html")
	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", struct {

		}{})

	})
}