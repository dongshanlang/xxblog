package repositories

import (
	"github.com/jinzhu/gorm"
	"xxblog/model"
)

var ArticleTagRepository = newArticleTagRepository()

type articleTagRepository struct {
}

func newArticleTagRepository() *articleTagRepository {
	return &articleTagRepository{}
}
func (r *articleTagRepository) Create(db *gorm.DB, t *model.ArticleTag) (err error) {
	err = db.Create(t).Error
	return
}
func (r *articleTagRepository) FindByArticleID(db *gorm.DB, Id int64) (at model.ArticleTag, err error) {
	err = db.Where("article_id = ?", Id).First(&at).Error
	return
}
