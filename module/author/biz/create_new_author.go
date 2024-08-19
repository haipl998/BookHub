package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"strings"

	"context"
)

type CreateAuthorStorage interface {
	GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Author, err error)
	CreateAuthor(ctx context.Context, data *model.Author) (err error)
}

type createAuthorBiz struct {
	store CreateAuthorStorage
}

func NewCreateAuthorBiz(store CreateAuthorStorage) *createAuthorBiz {
	return &createAuthorBiz{store: store}
}

func (biz *createAuthorBiz) CreateAuthor(ctx context.Context, data *model.Author) (err error) {
	err = checkBlankAuthor(data)
	if err != nil {
		common.ErrCannotCreateEntity(model.EntityName, err)
	}

	if _, err := biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.FirstName": data.FirstName, "Authors.LastName": data.LastName}); err == nil {
		return common.ErrEntityExisted(model.EntityName, common.EntityExisted)
	}

	err = biz.store.CreateAuthor(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}

// TODO: tìm các viết bằng method
func checkBlankAuthor(data *model.Author) error {
	data.FirstName = strings.TrimSpace(data.FirstName)
	if data.FirstName == "" {
		return model.ErrFirstNameIsBlank
	}

	data.LastName = strings.TrimSpace(data.LastName)
	if data.LastName == "" {
		return model.ErrLastNameIsBlank
	}

	return nil
}
