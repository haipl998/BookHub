package biz_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"
	"net/http"
	"time"

	"context"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type LoginStorage interface {
	GetMemberByEmail(ctx context.Context, cond map[string]interface{}) (result *model_member.SessionMember, err error)
}

type loginBiz struct {
	store LoginStorage
}

func NewLoginBiz(store LoginStorage) *loginBiz {
	return &loginBiz{store: store}
}

func (biz *loginBiz) Login(ctx context.Context, data *model_member.LoginForm) (cookie *http.Cookie, err error) {
	result, err := biz.store.GetMemberByEmail(ctx, map[string]interface{}{"Email": data.Email})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model_member.EntityName, err)
	}

	// So sánh mật khẩu
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password))
	if err != nil {
		return nil, common.ErrCannotLogin(err)
	}

	// Tạo JWT token
	token, err := GenerateJWT(*result)
	if err != nil {
		return nil, common.ErrCannotLogin(err)
	}

	// Thiết lập cookie chứa JWT token
	cookie = &http.Cookie{
		Name:     "token",                        // Tên của cookie
		Value:    token,                          // Giá trị của cookie là JWT token
		Expires:  time.Now().Add(24 * time.Hour), // Thời gian hết hạn của cookie (24 giờ)
		HttpOnly: true,                           // Bảo vệ cookie không bị truy cập thông qua JavaScript
		Secure:   true,                           // Cookie chỉ được gửi qua HTTPS
		Path:     "/",                            // Đường dẫn mà cookie có hiệu lực
	}

	return cookie, nil
}

func GenerateJWT(member model_member.SessionMember) (string, error) {
	// Tạo claims
	claims := model_member.Claims{
		MemberID: member.MemberID,
		Email:    member.Email,
		Role:     member.Role, // Vai trò của người dùng
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Thời hạn của JWT
			Issuer:    "your_app_name",
		},
	}

	// Tạo token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(common.JwtSecret)
}
