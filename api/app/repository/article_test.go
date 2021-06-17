package repository

import (
	"baf/api/domain"
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var a = &domain.Article{
	ID:    5,
	Title:  "Momo",
	Content: "momo@mail.com",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFind(t *testing.T) {
	db, mock := NewMock()
	repo := &articleRepo{db: db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, title, content FROM article"

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).AddRow(a.ID, a.Title, a.Content)
	mock.ExpectQuery(query).WithArgs(a.ID).WillReturnRows(rows)

	article, err := repo.FindArticleById(a.ID)
	assert.NotNil(t, article)
	assert.NoError(t, err)
}

func TestFindByIdErr(t *testing.T) {
	db, mock := NewMock()
	repo := &articleRepo{db: db}
	defer func() {
		repo.Close()
	}()

	query := "SELECT id, title, content FROM article WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "title", "content"})

	mock.ExpectQuery(query).WithArgs(a.ID).WillReturnRows(rows)

	article, err := repo.FindArticleById(a.ID)
	assert.Empty(t, article)
	assert.Error(t, err)
}
