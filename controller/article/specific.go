package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxblog/service"
)

type SpecificArticleReq struct {
	Id int64 `form:"id"`
}

func ShowSpecificArticle(ctx *gin.Context) {
	var req = SpecificArticleReq{}
	var err error
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	ctx.HTML(http.StatusOK, "content.html", gin.H{
		"article": service.ArticleService.GetArticle(req.Id),
		//"errmsg": "adfa",
	})
}
func DelArticle(ctx *gin.Context) {
	var req = SpecificArticleReq{}
	var err error
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	service.ArticleService.DelArticle(req.Id)
	ctx.Redirect(http.StatusFound, "/article/articles")
}
