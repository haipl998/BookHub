package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// gọi hàm đệ quy mục đich để lấy lỗi trong cùng của nó //Error(Error(Error)))
func (e *AppError) RootError() error {
	// Kiểm tra xem e.RootErr có phải là *AppError hay không
	if appErr, ok := e.RootErr.(*AppError); ok {
		return appErr.RootError()
	}
	// Nếu không, trả về e.RootErr như là lỗi gốc
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrorInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), "ErrorInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot list %s", strings.ToLower(entity)), fmt.Sprintf("ErrorCannotList%s", entity))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)), fmt.Sprintf("ErrorCannotDelete%s", entity))
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot update %s", strings.ToLower(entity)), fmt.Sprintf("ErrorCannotUpdate%s", entity))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot get %s", strings.ToLower(entity)), fmt.Sprintf("ErrorCannotGet%s", entity))
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s deleted", strings.ToLower(entity)), fmt.Sprintf("Error%sDeleted", entity))
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s already exists", strings.ToLower(entity)), fmt.Sprintf("Error%sAlreadyExists", entity))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s not found", strings.ToLower(entity)), fmt.Sprintf("Error%sNotFound", entity))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrorCannotCreate%s", entity))
}

func ErrCannotLogin(err error) *AppError {
	return NewCustomError(err, "you can not login", "ErrorCannotLogin")
}

func ErrPermission(err error) *AppError {
	return NewCustomError(err, "You have no permission", "ErrorNoPermission")
}

var RecordNotFound = errors.New("record not found")
