package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	 log "github.com/sirupsen/logrus"
	"net/http"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil {
		fmt.Println("token verified")
	} else {
		fmt.Println("masuk")
		result := gin.H{
			"code": 401,
			"message": "Unauthorized",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

func Auth2(c *gin.Context)  {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "root" && password == "root" {
		fmt.Println("called")
		log.WithFields(log.Fields{
			"user": user,
		}).Info("User authenticated")
	} else {
		fmt.Println("called 2")
		c.JSON(400, gin.H{"code" : 400, "message" : "Username or Password is incorrect"})
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		return
	}
}
