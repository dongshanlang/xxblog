package repositories

import (
	"github.com/jinzhu/gorm"
	"xxblog/model"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r *userRepository) Create(db *gorm.DB, t *model.User) (err error) {
	err = db.Create(t).Error
	return
}
func (r userRepository) FindByUsername(db *gorm.DB, t *model.User) (user model.User, err error) {
	err = db.Where("username = ?", t.Username).First(&user).Error
	return
}
