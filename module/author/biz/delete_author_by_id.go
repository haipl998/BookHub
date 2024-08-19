package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"

	"context"
)

type DeleteAuthorStorage interface {
	GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Author, err error)
	DeleteAuthorById(ctx context.Context, cond map[string]interface{}) (err error)
}

type deleteAuthorBiz struct {
	store DeleteAuthorStorage
}

func NewDeleteAuthorBiz(store DeleteAuthorStorage) *deleteAuthorBiz {
	return &deleteAuthorBiz{store: store}
}

func (biz *deleteAuthorBiz) DeleteAuthorById(ctx context.Context, id int) (err error) {
	_, err = biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	err = biz.store.DeleteAuthorById(ctx, map[string]interface{}{"AuthorID": id})
	if err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}
