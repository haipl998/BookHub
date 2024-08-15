package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"
	"strings"

	"context"
)

type CreateAuthorStorage interface {
	CreateAuthor(ctx context.Context, data *model.Authors) (err error)
}

type createAuthorBiz struct {
	store CreateAuthorStorage
}

func NewCreateAuthorBiz(store CreateAuthorStorage) *createAuthorBiz {
	return &createAuthorBiz{store: store}
}

func (biz *createAuthorBiz) CreateAuthor(ctx context.Context, data *model.Authors) (err error) {
	data.FirstName = strings.TrimSpace(data.FirstName)
	if data.FirstName == "" {
		return common.ErrorCannotCreateEntity(model.EntityName, model.ErrFirstNameIsBlank)
	}

	data.LastName = strings.TrimSpace(data.LastName)
	if data.LastName == "" {
		return common.ErrorCannotCreateEntity(model.EntityName, model.ErrLastNameIsBlank)
	}

	err = biz.store.CreateAuthor(ctx, data)
	if err != nil {
		if errors.Is(err, common.EntityExisted) {
			return common.ErrorEntityExisted(model.EntityName, err)
		}
		return common.ErrorCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
