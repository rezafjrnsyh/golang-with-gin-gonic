package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
