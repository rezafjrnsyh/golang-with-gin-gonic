package article

type Article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

type IArticleRepository interface {
	FindArticle() ([]*Article, error)
	CreateArticle(article *Article) (*Article, error)
	FindArticleById(id int) (*Article, error)
	Delete(id int) (int64,error)
}



