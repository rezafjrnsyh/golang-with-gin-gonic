package article

type Article struct {
	ID      int    `json:"id" form:"id"`
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
}



