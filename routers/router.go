package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxblog/controller/auth"
)

func Init(r *gin.Engine) {
	r.Static("static", "./static")
	r.LoadHTMLGlob("./views/*")
	//r.LoadHTMLFiles("views/login.html")

	//auth
	authRouter := r.Group("auth")
	{
		authRouter.GET("signin", auth.GetSignIn)
		authRouter.POST("signin", auth.SignIn)
		authRouter.GET("signup", auth.GetSignUp)
		authRouter.POST("signup", auth.SignUp)
	}

	test:=r.Group("/test")
	{
		test.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong")
		})
	}

}
func Compare(){

}