package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShowSpecificArticleReq struct {
	Id int64 `form:"id"`
}

func ShowSpecificArticle(ctx *gin.Context) {
	var req = ShowSpecificArticleReq{}
	var err error
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	ctx.HTML(http.StatusOK, "content.html", gin.H{})
}
