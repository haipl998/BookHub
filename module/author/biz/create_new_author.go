package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"

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
	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
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
