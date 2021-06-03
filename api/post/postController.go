package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var a = "test"
type name struct {
	success string
}
type Controller struct {}
func getPost(c *gin.Context) {
	if a != "hello" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while fetching the accounts"})
	}

	var test = name{success:"yes"}
	c.JSON(http.StatusOK, test)

}
