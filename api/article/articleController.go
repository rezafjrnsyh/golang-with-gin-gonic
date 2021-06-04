package article

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type articleController struct {
	ArticleService IArticleService
}

func ConstructorArticleController(db *sql.DB) *articleController {
	return &articleController{ConstructorArticleService(db)}
}

func InitializeRoutesArticle(db *sql.DB, r *gin.Engine)  {
	Controller := ConstructorArticleController(db)
	postRoutes := r.Group("/article")
		{
			//postRoutes.GET("/hello", GetPost)
			postRoutes.GET("/", Controller.GetAllArticle() )
		}
}


//func GetPost(c *gin.Context) {
//	if a != "hello" {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the accounts"})
//	}
//
//	//var test = name{success:"yes"}
//	c.String(http.StatusOK, "Hello golang with gin")
//
//}

func (s *articleController) GetAllArticle() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		Articles, err := s.ArticleService.GetArticles()
		fmt.Print("err",err)
		if err != nil {
			c.Writer.WriteHeader(http.StatusBadRequest)
			c.JSONP(http.StatusBadRequest, err.Error())
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.JSONP(http.StatusOK, Articles)
	}
}
