package service

import (
	"time"
	"xxblog/base/logger"
	"xxblog/model"
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
		artiTag, err := repositories.ArticleTagRepository.FindByArticleID(repositories.DB, a.Id)
		if err != nil {
			logger.Errorf("get article type by article id failed: %+v", err)
			continue
		}
		repositories.ArticleRepository
		articles = append(articles, Article{
			Id:       a.Id,
			ArtiName: a.Title,
			Atime:    time.Unix(a.CreateTime, 0),
			Acount:   a.ViewCount,
			ArticleType: {
				Id:    0,
				Tname: "",
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
func (s *articleService) AddArticle(title, content, imgUrl string, userId, articleType int64) bool {
	tx := repositories.DB.Begin()
	article := model.Article{
		UserId:      userId,
		Title:       title,
		Summary:     "",
		Content:     content,
		ContentType: "text",
		Status:      0,
		Share:       false,
		SourceUrl:   imgUrl,
		ViewCount:   0,
		CreateTime:  time.Now().Unix(),
		UpdateTime:  0,
	}
	err := repositories.ArticleRepository.Create(
		tx, &article)
	if err != nil {
		tx.Rollback()
		logger.Errorf("create article failed: %+v", err)
		return false
	}
	at := model.ArticleTag{
		ArticleId:  article.Id,
		TagId:      articleType,
		Status:     0,
		CreateTime: time.Now().Unix(),
	}
	err = repositories.ArticleTagRepository.Create(tx, &at)
	if err != nil {
		tx.Rollback()
		logger.Errorf("create article_tag failed: %+v", err)
		return false
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		logger.Errorf("create article commit failed: %+v", err)
		return false
	}
	return true
}

type ArticleShowInfo struct {
	ArtiName string
	TypeName string
	Acontent string
	Aimg     string
	Acount   int64
	Atime    time.Time
}

func (s *articleService) GetArticle(Id int64) *ArticleShowInfo {
	article, err := repositories.ArticleRepository.Get(repositories.DB, Id)
	if err != nil {
		logger.Errorf("get article error: %+v", err)
		return nil
	}
	return &ArticleShowInfo{
		ArtiName: article.Title,
		TypeName: "123",
		Acontent: article.Content,
		Aimg:     article.SourceUrl,
		Acount:   article.ViewCount,
		Atime:    time.Unix(article.CreateTime, 0),
	}
}
func (s *articleService) DelArticle(Id int64) {
	err := repositories.ArticleRepository.Del(repositories.DB, &model.Article{
		Model: model.Model{
			Id: Id,
		},
	})
	if err != nil {
		logger.Errorf("get article error: %+v", err)
		return
	}
	return
}
