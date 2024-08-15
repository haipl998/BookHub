package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

type ListAuthorStorage interface {
	ListAuthor(ctx context.Context) (result []model.Authors, err error)
}

type listAuthorBiz struct {
	store ListAuthorStorage
}

func NewListAuthorBiz(store ListAuthorStorage) *listAuthorBiz {
	return &listAuthorBiz{store: store}
}

func (biz *listAuthorBiz) ListAuthor(ctx context.Context) (result []model.Authors, err error) {
	result, err = biz.store.ListAuthor(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return result, nil

}
