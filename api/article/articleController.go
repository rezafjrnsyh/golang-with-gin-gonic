package article

import (
	"baf/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type articleController struct {
	ArticleService IArticleService
}

func CreateArticleController(db *sql.DB, r *gin.Engine)  {
	Controller := articleController{ArticleService: ConstructorArticleService(db)}
	postRoutes := r.Group("/api/article/")
		{
			postRoutes.GET("/list", utils.Auth, Controller.GetAllArticle )
			postRoutes.POST("/", utils.Auth, Controller.AddArticle)
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
		s := strings.Split(err.Error(), "'")
		errField := fmt.Errorf("field %s can't be empty", s[3])
		c.JSON(http.StatusBadRequest, gin.H{"message": errField.Error(), "code": 400})
	} else {
		title := c.DefaultPostForm("title", "Guest")
		content := c.PostForm("content")
		fmt.Println("controller : ", title , content)

		newArticle, err := s.ArticleService.AddArticle(&article)

		if err != nil {
			c.JSON(http.StatusInternalServerError,
				gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error"})
			//return
		} else {
			c.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "Article successfully created", newArticle))
		}
	}
}
