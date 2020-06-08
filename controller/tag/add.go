package tag

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xxblog/service"
)

func AddTag(ctx *gin.Context) {
	tags := service.TagService.GetAllTags()
	ctx.HTML(http.StatusOK, "addType.html", gin.H{
		"articleType": tags,
	})
}
