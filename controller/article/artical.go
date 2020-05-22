package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxblog/service"
)

//
type ArticleReq struct {
	Select int64 `form:"select"`
}

func ShowArticles(ctx *gin.Context) {
	pagination := service.ArticleService.GetPagination()
	var req = ArticleReq{}
	var selectNum int64
	var err error
	err = ctx.ShouldBind(&req)
	if err != nil || req.Select == 0 {
		selectNum = 4
	} else {
		selectNum = req.Select
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"articleType": service.ArticleService.GetArticleTypes(),
		"title":       pagination.PageTitle,
		"articles":    service.ArticleService.GetArticles(),
		"pageIndex":   pagination.PageIndex,
		"pageCount":   pagination.PageCount,
		"count":       pagination.Count,
		"typeid":      selectNum,
	})
}
func ShowAddArticle(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add.html", gin.H{
		"articleType": service.ArticleService.GetArticleTypes(),
	})
}
func AddArticle(ctx *gin.Context) {

}
