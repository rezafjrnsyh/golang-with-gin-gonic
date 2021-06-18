package controller

import (
	"baf/api/app/domain"
	"baf/api/app/service"
	"baf/middleware"
	"baf/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type articleController struct {
	ArticleService domain.IArticleService
}

const (
	ARTICLE_LIST_PATH  = "/article/list"
	ARTICLE_CREATE_PATH = "/article"
	ARTICLE_GET_BY_ID_PATH = "/article/:id"
	ARTICLE_DELETE_PATH = "/article/:id"
)

func NewArticleController(db *sql.DB, r *gin.RouterGroup)  {
	Controller := articleController{ArticleService: service.NewArticleService(db)}
	r.GET(ARTICLE_LIST_PATH, middleware.Auth2, Controller.GetAllArticle )
	r.POST(ARTICLE_CREATE_PATH, middleware.Auth2, Controller.AddArticle)
	r.GET(ARTICLE_GET_BY_ID_PATH, middleware.Auth2, Controller.GetArticleById)
	r.DELETE(ARTICLE_DELETE_PATH, middleware.Auth2, Controller.DeleteArticle)
}

func (s *articleController) GetAllArticle(c *gin.Context) {
	articles, err := s.ArticleService.GetArticles()
	fmt.Print("err",err)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", articles))
}

func (s *articleController) AddArticle(c *gin.Context) {
	var article domain.Article
	err := c.BindJSON(&article)
	if err != nil {
		s := strings.Split(err.Error(), "'")
		errField := fmt.Errorf("field %s can't be empty", s[3])
		c.JSON(http.StatusBadRequest, gin.H{"message": errField.Error(), "code": 400})
	} else {
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

func (s *articleController) GetArticleById(c *gin.Context) {
	param := c.Param("id")
	id,err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, gin.H{"code" : 500, "message" : "Internal Server Error"})
	}
	article, er := s.ArticleService.GetArticle(id)
	if er != nil {
		log.Println(er.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code" : 400, "message" : "data not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "ok", "data": article})
	}
}

func (s *articleController) DeleteArticle(c *gin.Context) {
	param := c.Param("id")
	id,err := strconv.Atoi(param)
	if err != nil {
		log.Println("Failed to converted to int")
		c.JSON(http.StatusInternalServerError, gin.H{"code" : 500, "message" : "Internal Server Error"})
	}
	result, err := s.ArticleService.DeleteArticle(id)
	log.Println("rows:",result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code" : 500, "message" : "Internal server error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Data deleted successfully", "data": result})
	}
}
