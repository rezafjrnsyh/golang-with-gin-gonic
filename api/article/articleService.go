package article

import "database/sql"

type articleService struct {
	db	*sql.DB
	ArticleRepo IArticleRepository
}

type IArticleService interface {
	GetArticles() (*[]Article, error)
}

func ConstructorArticleService(db *sql.DB) IArticleService {
	return &articleService{db, NewArticleRepo(db)}
}

func (a *articleService) GetArticles() (*[]Article, error) {
	articles,err := a.ArticleRepo.GetAllArticle()
	if err != nil {
		return nil, err
	}

	return articles, nil
}


