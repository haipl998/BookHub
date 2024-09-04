package common

import (
	"github.com/dgrijalva/jwt-go"
)

var (
	JwtSecret = []byte("your_secret_key")
)

type Claims struct {
	MemberID int    `json:"member_id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
