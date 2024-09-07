package middleware

import (
	"BookHub/common"
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phoneRegex = `^\+?[0-9]\d{1,14}$`
)

func ValidateEmailAndPhone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input map[string]interface{}
		if err := c.ShouldBindBodyWith(&input, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			c.Abort()
			return
		}

		// Validate Email
		email, ok := input["Email"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(errors.New("email is required")))
			c.Abort()
			return
		}
		if !regexp.MustCompile(emailRegex).MatchString(email) {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(errors.New("invalid email format")))
			c.Abort()
			return
		}

		// Validate Phone
		phoneNumber, ok := input["PhoneNumber"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(errors.New("phone number is required")))
			c.Abort()
			return
		}
		if !regexp.MustCompile(phoneRegex).MatchString(phoneNumber) {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(errors.New("invalid phone number format")))
			c.Abort()
			return
		}

		c.Next()
	}
}
