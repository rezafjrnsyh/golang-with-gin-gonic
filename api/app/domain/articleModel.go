package domain

type Article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type IArticleService interface {
	GetArticles() ([]*Article, error)
	AddArticle(article *Article) (*Article, error)
	GetArticle(id int) (*Article, error)
	DeleteArticle(id int) (int64,error)
}

type IArticleRepository interface {
	FindArticle() ([]*Article, error)
	CreateArticle(article *Article) (*Article, error)
	FindArticleById(id int) (*Article, error)
	Update(article *Article) (*Article, error)
	Delete(id int) (int64,error)
}



