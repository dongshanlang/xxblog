package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxblog/model"
	"xxblog/service"
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

func GetSignIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", struct{}{})
}
func SignIn(ctx *gin.Context) {
	var (
		err error
		req SignInReq
	)
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", struct{}{})
		return
	}
	err = service.UserService.SignIn(&model.User{
		Username: sql.NullString{
			String: req.UserName,
			Valid:  len(req.UserName) > 0,
		},
		Password: req.Password,
	})
	if err != nil {
		ctx.HTML(http.StatusOK, "login.html", struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func GetSignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", struct{}{})
}

func SignUp(ctx *gin.Context) {
	var (
		err error
		req SignUpReq
	)
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.HTML(http.StatusOK, "register.html", struct{}{})
		return
	}
	err = service.UserService.SignUp(&model.User{
		Username: sql.NullString{
			String: req.UserName,
			Valid:  len(req.UserName) > 0,
		},
		Password: req.Password,
	})
	if err != nil {
		ctx.HTML(http.StatusOK, "register.html", struct{}{})
		return
	}
	ctx.Redirect(http.StatusFound, "/auth/signin")
}
