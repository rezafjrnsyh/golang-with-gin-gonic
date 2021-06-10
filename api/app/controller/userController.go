package controller

import (
	"baf/api/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func CreateUserController(r *gin.Engine) {
	userRoutes := r.Group("/api/user")
	{
		userRoutes.POST("/login", loginHandler)
	}
}

func loginHandler(c *gin.Context) {
	var credential domain.Credential
	err := c.Bind(&credential)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if credential.Username != domain.User1.Username || credential.Password != domain.User1.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Please provide valid login details",
		})
		return
	}

	token, err := CreateToken(domain.User1.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CreateToken(userId uint64) (string, error) {
	var err error
	//Creating Access Token
	_ = os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
