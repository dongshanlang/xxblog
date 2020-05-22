package service

import (
	"github.com/jinzhu/gorm"
	"xxblog/base/logger"
	"xxblog/model"
	"xxblog/repositories"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

//登陆
func (s *userService) SignIn(user *model.User) (err error) {
	u, err := repositories.UserRepository.FindByUsername(repositories.DB, user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Infof("no user found: %+v", user)
			return err
		}
		logger.Errorf("get user info error: %+v", err)
		return err
	}
	if u.Password == user.Password {
		return nil
	}
	return
}

//注册
func (s *userService) SignUp(user *model.User) (err error) {
	err = repositories.UserRepository.Create(repositories.DB, user)
	if err != nil {
		logger.Errorf("sign up failed: %+v", err)
		return
	}
	return
}
