package routers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"xxblog/controller/article"
	"xxblog/controller/auth"
	"xxblog/controller/tag"
)

func Init(r *gin.Engine) {

	//auth
	authRouter := r.Group("auth")
	{
		authRouter.GET("signin", auth.GetSignIn)
		authRouter.POST("signin", auth.SignIn)
		authRouter.GET("signup", auth.GetSignUp)
		authRouter.POST("signup", auth.SignUp)
	}
	articleRoute := r.Group("article")
	{
		articleRoute.GET("articles", article.ShowArticles)
		articleRoute.GET("add", article.ShowAddArticle)
		articleRoute.POST("add", article.AddArticle)
		articleRoute.GET("article", article.ShowSpecificArticle)
		articleRoute.GET("del", article.DelArticle)
	}
	tagRoute := r.Group("tag")
	{
		tagRoute.GET("add", tag.AddTag)
		tagRoute.POST("add", tag.AddTag)
	}

	test := r.Group("/test")
	{
		test.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong")
		})
	}
	r.SetFuncMap(template.FuncMap{
		"showprepage":  prepage,
		"shownextpage": shownextpage,
		"compare":      compare,
	})
	r.Static("static", "./static")
	r.LoadHTMLGlob("./views/*")
}

//试图函数，获取上一页页码

/*
1.在试图中定义视图函数函数名    | funcName

2.一般在main.go里面实现试图函数

3.在main函数里面把实现的函数和试图函关联起来   beego.AddFuncMap()
*/
func prepage(pageindex int) (preIndex int) {
	preIndex = pageindex - 1
	return
}

func shownextpage(pageindex int) (nextIndex int) {
	nextIndex = pageindex + 1
	return
}
func compare(a, b interface{}) bool {
	if a == nil || b == nil {
		return false
	}
	return a == b
}
