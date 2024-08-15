package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"
	"strings"

	"context"
)

type UpdateAuthorStorage interface {
	GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Authors, err error)
	UpdateAuthor(ctx context.Context, data *model.AuthorUpdate, cond map[string]interface{}) (err error)
}

type updateAuthorBiz struct {
	store UpdateAuthorStorage
}

func NewUpdateAuthorBiz(store UpdateAuthorStorage) *updateAuthorBiz {
	return &updateAuthorBiz{store: store}
}

func (biz *updateAuthorBiz) UpdateAuthor(ctx context.Context, data *model.AuthorUpdate, id int) (err error) {
	data.FirstName = strings.TrimSpace(data.FirstName)
	data.LastName = strings.TrimSpace(data.LastName)
	if data.LastName == "" && data.FirstName == "" {
		return common.ErrorCannotUpdaterEntity(model.EntityName, model.ErrBothIsBlank)
	}

	_, err = biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrorCannotGetEntity(model.EntityName, err)
		}
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}

	err = biz.store.UpdateAuthor(ctx, data, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		return common.ErrorCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
