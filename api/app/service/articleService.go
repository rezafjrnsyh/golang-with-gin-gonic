package service

import (
	"baf/api/app/domain"
	"baf/api/app/repository"
	"database/sql"
	"fmt"
)

type articleService struct {
	db          *sql.DB
	ArticleRepo domain.IArticleRepository
}

func NewArticleService(db *sql.DB) domain.IArticleService {
	return &articleService{db, repository.NewArticleRepo(db)}
}

func (a *articleService) GetArticles() ([]*domain.Article, error) {
	articles,err := a.ArticleRepo.FindArticle()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *articleService) AddArticle(article *domain.Article) (*domain.Article, error) {
	article, err := a.ArticleRepo.CreateArticle(article)
	fmt.Println("Service :", article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *articleService) GetArticle(id int) (*domain.Article, error) {
	article, err := a.ArticleRepo.FindArticleById(id)
	if err != nil{
		return nil, err
	}

	return article, nil
}

func (a *articleService) DeleteArticle(id int) (int64,error) {
	article, err := a.ArticleRepo.FindArticleById(id)
	if err != nil{
		return 0, err
	}

	result, err := a.ArticleRepo.Delete(article.ID)
	if err != nil {
		return result, err
	}
	return result, nil


}


