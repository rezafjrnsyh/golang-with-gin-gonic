package article

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type articleRepo struct {
	db	*sql.DB
}

type RepositoryArticle interface {
	GetAllArticle() (*[]Article, error)
}

func ConstructorArticle(db *sql.DB) RepositoryArticle {
	return &articleRepo{db}
}

func (a *articleRepo) GetAllArticle() (*[]Article, error) {
	var article Article
	var articles []Article

	queryInput := fmt.Sprintf("SELECT * FROM article")
	rows, err := a.db.Query(queryInput)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&article.ID, &article.Title, &article.Content)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		articles = append(articles, article)
	}
	return &articles, nil
}
