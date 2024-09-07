package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"

	"context"
)

type ListCategoryStorage interface {
	ListCategory(ctx context.Context) (result *[]model.Category, err error)
}

type listCategoryBiz struct {
	store ListCategoryStorage
}

func NewListCategoryBiz(store ListCategoryStorage) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context) (result *[]model.Category, err error) {
	result, err = biz.store.ListCategory(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return result, nil

}
