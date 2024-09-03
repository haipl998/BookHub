package biz_member

import (
	"context"
	"net/http"
	"time"
)

type logoutBiz struct {
}

func NewLogoutBiz() *logoutBiz {
	return &logoutBiz{}
}

func (biz *logoutBiz) Logout(ctx context.Context) (cookie *http.Cookie) {
	// Xóa cookie bằng cách thiết lập giá trị rỗng và thời gian hết hạn là một thời điểm trong quá khứ
	cookie = &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Đặt thời gian hết hạn là một giờ trước
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	return cookie

}
