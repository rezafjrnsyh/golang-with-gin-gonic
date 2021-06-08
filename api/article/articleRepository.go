package article

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type articleRepo struct {
	db	*sql.DB
}

type IArticleRepository interface {
	GetAllArticle() ([]Article, error)
	AddArticle(article *Article) (*Article, error)
}

func NewArticleRepo(db *sql.DB) IArticleRepository {
	return &articleRepo{db}
}

func (a *articleRepo) GetAllArticle() ([]Article, error) {
	var article Article
	var articles []Article

	queryInput := fmt.Sprintf("SELECT * FROM article")
	rows, err := a.db.Query(queryInput)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&article.ID,&article.Title, &article.Content)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *articleRepo) AddArticle(article *Article) (*Article, error) {
	query := fmt.Sprintf(`INSERT INTO article(id,title, content) VALUES (?, ?,?)`)
	stmnt, err := a.db.Prepare(query)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}
	defer stmnt.Close()

	result, err := stmnt.Exec(&article.ID, &article.Title, &article.Content)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}

	lastIndex, _ := result.LastInsertId()

	return a.GETMenu(int(lastIndex))
}

func (p *articleRepo) GETMenu(id int) (*Article, error) {
	fmt.Println(id)
	results := p.db.QueryRow("SELECT * FROM article WHERE id=?", id)

	var a Article
	err := results.Scan(&a.ID,&a.Title, &a.Content)
	if err != nil {
		return nil, errors.New("Article ID Not Found")
	}

	return &a, nil
}
