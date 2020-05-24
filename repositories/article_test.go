package repositories

import (
	"testing"
	"time"
	"xxblog/model"
)

func TestCreateArticle(t *testing.T) {
	tx := DB.Begin()
	article := model.Article{
		UserId:      1,
		Title:       "体育新闻",
		Summary:     "新闻呗",
		Content:     "科比挂了，唉。。",
		ContentType: "text",
		Status:      0,
		Share:       false,
		SourceUrl:   "",
		ViewCount:   1,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  0,
	}
	err := ArticleRepository.Create(
		tx, &article)
	if err != nil {
		tx.Rollback()
		t.Log(err)
		return
	}
	at := model.ArticleTag{
		ArticleId:  article.Id,
		TagId:      1,
		Status:     0,
		CreateTime: time.Now().Unix(),
	}
	err = ArticleTagRepository.Create(tx, &at)
	if err != nil {
		tx.Rollback()
		t.Log(err)
		return
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		t.Log(err)
		return
	}
}
