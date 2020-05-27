package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xxblog/base/logger"
	"xxblog/service"
)

type AddArticleReq struct {
	ArticleName string `form:"articleName"`
	Select      int64  `form:"select"`
	Content     string `form:"content"`
	//FileInfo    os.File `form:"uploadname"`
}

func AddArticle(ctx *gin.Context) {
	var req = AddArticleReq{}
	var err error
	var dst string
	var bCreateRet bool
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}
	file, err := ctx.FormFile("uploadname")
	if err == nil && file != nil {
		// Upload the file to specific dst.
		dst = fmt.Sprintf("./static/upload/%s", file.Filename)
		err = ctx.SaveUploadedFile(file, dst)
		if err != nil {
			ctx.JSON(http.StatusOK, err)
			return
		}
	} else if err != nil {
		logger.Infof("get file failed: %+v", err)
		err = nil
	}
	if dst == "" {
		bCreateRet = service.ArticleService.AddArticle(req.ArticleName, req.Content, "", 1, req.Select)
	} else {
		bCreateRet = service.ArticleService.AddArticle(req.ArticleName, req.Content, dst[1:], 1, req.Select)
	}

	if !bCreateRet {
		ctx.HTML(http.StatusOK, "add.html", gin.H{
			"articleType": service.ArticleService.GetArticleTypes(),
			"errmsg":      "文章添加失败",
		})
	}
	ctx.Redirect(http.StatusFound, "/article/articles")
}
