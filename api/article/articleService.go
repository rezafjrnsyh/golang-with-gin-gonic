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
	GetArticles() ([]*Article, error)
	AddArticle(article *Article) (*Article, error)
	GetArticle(id int) (*Article, error)
	DeleteArticle(id int) (int64,error)
}

func ConstructorArticleService(db *sql.DB) IArticleService {
	return &articleService{db, NewArticleRepo(db)}
}

func (a *articleService) GetArticles() ([]*Article, error) {
	articles,err := a.ArticleRepo.FindArticle()
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (a *articleService) AddArticle(article *Article) (*Article, error) {
	article, err := a.ArticleRepo.CreateArticle(article)
	fmt.Println("Service :", article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a *articleService) GetArticle(id int) (*Article, error) {
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


