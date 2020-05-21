package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type SignInReq struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"pwd" form:"pwd"`
	Remember string `json:"remember" form:"remember"`
}
type SignUpReq struct {
	UserName string `json:"username" form:"username"`
	Password string `json:"pwd" form:"pwd"`
}

func GetSignIn(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "login.html", struct {
	}{})
}
func SignIn(ctx *gin.Context)  {
	var (
		err error
		req SignInReq
	)
	err = ctx.ShouldBind(&req)
	if err!=nil{
		ctx.JSON(http.StatusOK, err.Error())
	}
	 ctx.JSON(http.StatusOK, "ok")
}

func GetSignUp(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "register.html", struct {	}{})
}

func SignUp(ctx *gin.Context)  {
	var (
		err error
		req SignUpReq
	)
	err = ctx.ShouldBind(&req)
	if err!=nil{
		ctx.JSON(http.StatusOK, err.Error())
	}
	ctx.JSON(http.StatusOK, "ok")
}