package middleware

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"
	"context"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthenticateJWT là middleware để xác thực JWT
func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ cookie
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		tokenString := cookie.Value

		// Parse và kiểm tra token
		claims := &model_member.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Lưu thông tin người dùng vào context để sử dụng sau này
		ctx := context.WithValue(c.Request.Context(), "member", claims)
		c.Request = c.Request.WithContext(ctx)

		// Tiếp tục thực hiện request
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
		// Lấy cookie chứa token từ request
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token provided"})
			c.Abort()
			return
		}

		// Parse token và xác thực
		claims := &model_member.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
			c.Abort()
			return
		}

		// Kiểm tra role của user
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - You do not have access"})
			c.Abort()
			return
		}

		// Nếu là admin, cho phép tiếp tục request
		c.Next()
	}
}

func AuthorizeSelf() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - No token provided"})
			c.Abort()
			return
		}

		// Parse token và xác thực
		claims := &model_member.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(common.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid token"})
			c.Abort()
			return
		}

		// Lấy member ID từ URL parameter
		//memberID := c.Param("id")
		memberID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// Kiểm tra quyền truy cập
		if claims.Role == "admin" || claims.MemberID == memberID {
			// Admin hoặc chính member đó được phép truy cập
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden - You do not have access"})
			c.Abort()
		}
	}
}
