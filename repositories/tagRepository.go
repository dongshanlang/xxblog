package repositories

import (
	"github.com/jinzhu/gorm"
	"xxblog/model"
)

var TagRepository = newTagRepository()

type tagRepository struct {
}

func newTagRepository() *tagRepository {
	return &tagRepository{}
}
func (r *tagRepository) Create(db *gorm.DB, t *model.Tag) (err error) {
	err = db.Create(t).Error
	return
}
func (r *tagRepository) GetAll(db *gorm.DB) (tags []model.Tag, err error) {
	err = db.Where("status = ?", 0).Find(&tags).Error
	return
}
