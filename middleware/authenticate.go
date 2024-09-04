package middleware

import (
	"BookHub/common"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// getAndValidateToken lấy token từ cookie và xác thực nó
func getAndValidateToken(c *gin.Context) (*jwt.Token, *common.Claims, error) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		return nil, nil, err
	}

	claims := &common.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, nil, err
	}

	return token, claims, nil
}

// AuthenticateJWT là middleware để xác thực JWT
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, claims, err := getAndValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("member", claims)
		c.Next()
	}
}

// Lấy thông tin người dùng từ context
// func GetUserFromContext(c *gin.Context) (*model_member.Claims, bool) {
// 	user, exists := c.Request.Context().Value("user").(*model_member.Claims)
// 	return user, exists
// }

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, claims, err := getAndValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - You do not have access"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AuthorizeSelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, claims, err := getAndValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		memberID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if claims.Role == "admin" || claims.MemberID == memberID {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - You do not have access"})
			c.Abort()
		}
	}
}
