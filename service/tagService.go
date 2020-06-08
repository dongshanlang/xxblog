package service

import (
	"time"
	"xxblog/base/logger"
	"xxblog/model"
	"xxblog/repositories"
)

var TagService = newTagService()

func newTagService() *tagService {
	return &tagService{}
}

type tagService struct {
}

//{{range .articleType}}
//<tr>
//<td>{{.Id}}</td>
//<td>{{.Tname}}</td>
type Tag struct {
	Id    int64
	Tname string
}

func (ts *tagService) GetAllTags() (tgs []Tag) {
	tags, err := repositories.TagRepository.GetAll(repositories.DB)
	if err != nil {
		logger.Errorf("get tags failed: %+v", err)
		return nil
	}
	for _, t := range tags {
		tgs = append(tgs, Tag{
			Id:    t.Id,
			Tname: t.Name,
		})
	}
	return
}

func (ts *tagService) Create(name string) bool {
	err := repositories.TagRepository.Create(repositories.DB, &model.Tag{
		//Model:       model.Model{},
		Name:        name,
		Description: "",
		Status:      0,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  time.Now().Unix(),
	})
	if err != nil {
		logger.Errorf("create tag failed: %+v", err)
		return false
	}
	return true
}
