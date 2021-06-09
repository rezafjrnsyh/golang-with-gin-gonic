package article

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)

type articleRepo struct {
	db	*sql.DB
}

func NewArticleRepo(db *sql.DB) IArticleRepository {
	return &articleRepo{db}
}

func (a *articleRepo) FindArticle() ([]*Article, error) {
	articles := make([]*Article, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := a.db.QueryContext(ctx, "SELECT id, title, content FROM article")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		article := new(Article)
		err := rows.Scan(&article.ID,&article.Title, &article.Content)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *articleRepo) CreateArticle(article *Article) (*Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf(`INSERT INTO article(id,title, content) VALUES (?, ?,?)`)
	stmt, err := a.db.PrepareContext(ctx,query)
	if err != nil {
		s := strings.Split(err.Error(), ":")
		log.Println(s[1])
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, &article.ID, &article.Title, &article.Content)
	fmt.Println("res:", &result)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}

	lastIndex, _ := result.LastInsertId()

	return a.FindArticleById(int(lastIndex))
}

func (a *articleRepo) FindArticleById(id int) (*Article, error) {
	article := new(Article)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := a.db.QueryRowContext(ctx, "SELECT id, title, content FROM article WHERE id=?", id).
		Scan(&article.ID,&article.Title, &article.Content)
	if err != nil {
		return nil, errors.New("article ID not found")
	}
	return article, nil
}

func (a *articleRepo) Delete(id int) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM article WHERE id = ?"
	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
		errorMsg := errors.New("DATABASE ERROR - " + err.Error())
		return 0, errorMsg
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected Error", err)
	}
	return RowsAffected,nil
}
