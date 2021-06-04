package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var a = "hello"
type name struct {
	success string
}

func InitializeRoutesPost(r *gin.Engine)  {
	postRoutes := r.Group("/post")
		{
			postRoutes.GET("/hello", GetPost)
		}
}

func GetPost(c *gin.Context) {
	if a != "hello" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the accounts"})
	}

	//var test = name{success:"yes"}
	c.String(http.StatusOK, "Hello golang with gin")

}
