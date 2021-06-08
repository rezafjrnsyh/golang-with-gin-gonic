package article

type Article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}

var Schema = `
DROP TABLE IF EXISTS article;
CREATE TABLE article (
	id    INTEGER PRIMARY KEY,
    title VARCHAR(80)  DEFAULT '',
    content  VARCHAR(330)  DEFAULT '',
);`



