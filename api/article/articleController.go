package article

import (
	"baf/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type articleController struct {
	ArticleService IArticleService
}

func CreateArticleController(db *sql.DB, r *gin.Engine)  {
	Controller := articleController{ArticleService: ConstructorArticleService(db)}
	postRoutes := r.Group("/api/article/")
		{
			postRoutes.GET("/list", Controller.GetAllArticle )
			postRoutes.POST("/", Controller.AddArticle)
		}
}

func (s *articleController) GetAllArticle(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	articles, err := s.ArticleService.GetArticles()
	fmt.Print("err",err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", articles))
}

func (s *articleController) AddArticle(c *gin.Context) {
	var article Article
	err := c.BindJSON(&article)
	if err != nil {
		fmt.Println(err["Error"].)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		title := c.DefaultPostForm("title", "Guest")
		content := c.PostForm("content")
		fmt.Println("controller : ", title , content)

		newArticle, err := s.ArticleService.AddArticle(&article)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "Article successfully created", newArticle))
	}
}
