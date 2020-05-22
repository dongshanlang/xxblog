package service

import (
	"time"
	"xxblog/repositories"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

type ArtiType struct {
	Id    int64
	Tname string
}

type Pagination struct {
	PageTitle string
	PageIndex int
	PageCount int
	Count     int
}

type Article struct {
	Id          int64
	ArtiName    string
	Atime       time.Time
	Acount      int64
	ArticleType ArtiType
}

func (s *articleService) GetArticleTypes() (types []ArtiType) {
	list, err := repositories.TagRepository.GetAll(repositories.DB)
	if err != nil {
		return nil
	}
	for _, tag := range list {
		types = append(types, ArtiType{
			Id:    tag.Id,
			Tname: tag.Name,
		})
	}
	return
}

func (s *articleService) GetArticles() (articles []Article) {
	list, err := repositories.ArticleRepository.GetUserAll(repositories.DB, 1)
	if err != nil {
		return nil
	}
	for _, a := range list {
		articles = append(articles, Article{
			Id:       a.Id,
			ArtiName: a.Title,
			Atime:    time.Unix(a.CreateTime, 0),
			Acount:   a.ViewCount,
			ArticleType: ArtiType{
				Id:    4,
				Tname: "体育新闻",
			},
		})
	}

	return
}
func (s *articleService) GetPagination() (pagination Pagination) {
	return Pagination{
		PageTitle: "main",
		PageIndex: 1,
		PageCount: 10,
		Count:     100,
	}
}
