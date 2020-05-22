package repositories

import (
	"os"
	"testing"
	"time"
	"xxblog/base/conf"
	"xxblog/model"
)

func TestMain(m *testing.M) {
	conf.Init()
	InitDBConnection(model.Models...)
	os.Exit(m.Run())
}
func TestAdd(t *testing.T) {
	tag := model.Tag{
		Name:        "体育新闻",
		Description: "体育类新闻",
		Status:      0,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  0,
	}
	err := TagRepository.Create(DB, &tag)
	if err != nil {
		t.Log(err)
		return
	}

	tag1 := model.Tag{
		Name:        "财经新闻",
		Description: "财经类新闻",
		Status:      0,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  0,
	}
	err = TagRepository.Create(DB, &tag1)
	if err != nil {
		t.Log(err)
		return
	}
}
