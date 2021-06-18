package repository

import (
	"baf/api/app/domain"
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
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

func TestArticleRepo_FindArticleById(t *testing.T) {
	db, mock := NewMock()
	r := &articleRepo{db: db}
	defer func() {
		r.Close()
	}()

	tests := []struct{
		name string
		r domain.IArticleRepository
		msgId int
		mock func()
		want *domain.Article
		wantErr bool
	}{
		{
			name:  "OK",
			r:     r,
			msgId: 1,
			mock: func() {
				//We added one row
				rows := sqlmock.NewRows([]string{"Id", "Title", "Content"}).AddRow(1, "title", "body")
				t.Log("logs rows: " , rows)
				query := "SELECT id, title, content FROM article"
				mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			want: &domain.Article{
				ID:        1,
				Title:     "title",
				Content:   "body",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Error("log")
			tt.mock()
			got, err := tt.r.FindArticleById(tt.msgId)
			t.Error(got)
			if (err != nil) != tt.wantErr {
				// t.Error("called")
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				// t.Error("called2")
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
