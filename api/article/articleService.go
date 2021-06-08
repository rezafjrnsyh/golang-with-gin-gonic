package article

import (
	"database/sql"
	"fmt"
)

type articleService struct {
	db	*sql.DB
	ArticleRepo IArticleRepository
}

type IArticleService interface {
	GetArticles() ([]Article, error)
	AddArticle(article *Article) (*Article, error)
}

func ConstructorArticleService(db *sql.DB) IArticleService {
	return &articleService{db, NewArticleRepo(db)}
}

func (a *articleService) GetArticles() ([]Article, error) {
	articles,err := a.ArticleRepo.GetAllArticle()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *articleService) AddArticle(article *Article) (*Article, error) {
	article, err := a.ArticleRepo.AddArticle(article)
	fmt.Println("Service :", article)
	if err != nil {
		return nil, err
	}

	return article, nil
}


