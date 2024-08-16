package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"errors"
	"strings"

	"context"
)

type CreateCategoryStorage interface {
	CreateCategory(ctx context.Context, data *model.Category) (err error)
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *model.Category) (err error) {
	data.CategoryName = strings.TrimSpace(data.CategoryName)
	if data.CategoryName == "" {
		return common.ErrorCannotCreateEntity(model.EntityName, model.ErrCategoryNameIsBlank)
	}

	if err = biz.store.CreateCategory(ctx, data); err != nil {
		if errors.Is(err, common.EntityExisted) {
			return common.ErrorEntityExisted(model.EntityName, err)
		}
		return common.ErrorCannotCreateEntity(model.EntityName, err)
	}
	return nil

}
