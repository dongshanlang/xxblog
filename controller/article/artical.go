package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func ShowArticles(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", "")
}
