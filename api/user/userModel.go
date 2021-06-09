package user

import "github.com/dgrijalva/jwt-go"

type Credential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type MyClaims struct {
		jwt.StandardClaims
		Username string `json:"Username"`
		Email    string `json:"Email"`
}

var User1 = User{
	ID:            1,
	Username: "root",
	Password: "root",
	Phone: "49123454322", //this is a random number
	Email: "test@gmail.com",
}
