package repository

import (
	"baf/api/domain"
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

func NewArticleRepo(db *sql.DB) domain.IArticleRepository {
	return &articleRepo{db:db}
}

func (a *articleRepo) FindArticle() ([]*domain.Article, error) {
	articles := make([]*domain.Article, 0)

	query := fmt.Sprintf(`SELECT id, title, content FROM article`)
	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		article := &domain.Article{}
		err := rows.Scan(&article.ID,&article.Title, &article.Content)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *articleRepo) CreateArticle(article *domain.Article) (*domain.Article, error) {

	query := fmt.Sprintf(`INSERT INTO article(id,title, content) VALUES (?, ?,?)`)
	result, err := a.db.Exec(query, &article.ID, &article.Title, &article.Content)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}

	id, _ := result.LastInsertId()

	return a.FindArticleById(int(id))
}

func (a *articleRepo) FindArticleById(id int) (*domain.Article, error) {
	article := new(domain.Article)

	query := fmt.Sprintf(`SELECT id, title, content FROM article WHERE id=?`)
	err := a.db.QueryRow(query, id).
		Scan(&article.ID,&article.Title, &article.Content)
	if err != nil {
		return nil, errors.New("article ID not found")
	}
	return article, nil
}

func (a *articleRepo) Delete(id int) (int64, error) {

	query := "DELETE FROM article WHERE id = ?"
	result, err := a.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected Error", err)
	}
	return RowsAffected,nil
}

func (a *articleRepo) Update(article *domain.Article) (*domain.Article, error) {
	query := "UPDATE article SET id = ?, title = ?, content = ? WHERE id = ?"
	result, err := a.db.Exec(query, &article.ID, &article.Title, &article.Content)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}
	lastIndex, _ := result.LastInsertId()

	return a.FindArticleById(int(lastIndex))
}
