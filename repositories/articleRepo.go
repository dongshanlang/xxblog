package repositories

import (
	"github.com/jinzhu/gorm"
	"xxblog/model"
)

var ArticleRepository = newArticleRepository()

type articleRepository struct {
}

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}
func (r *articleRepository) Create(db *gorm.DB, t *model.Article) (err error) {
	err = db.Create(t).Error
	return
}
func (r *articleRepository) Del(db *gorm.DB, t *model.Article) (err error) {
	err = db.Delete(t).Error
	return
}
func (r *articleRepository) GetUserAll(db *gorm.DB, userId int64) (list []model.Article, err error) {
	var articles []model.Article
	err = db.Where("user_id = ?", userId).Find(&articles).Error
	return articles, err
}
func (r *articleRepository) GetTitleAll(db *gorm.DB, title string) (list []model.Article, err error) {
	var articles []model.Article
	err = db.Where("title = ?", title).Find(&articles).Error
	return articles, err
}
func (r *articleRepository) Get(db *gorm.DB, Id int64) (article model.Article, err error) {
	err = db.Where("id = ?", Id).Find(&article).Error
	return article, err
}
