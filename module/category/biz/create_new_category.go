package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"strings"

	"context"
)

type CreateCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (result *model.Category, err error)
	CreateCategory(ctx context.Context, data *model.Category) (err error)
}

type createCategoryBiz struct {
	store CreateCategoryStorage
}

func NewCreateCategoryBiz(store CreateCategoryStorage) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *model.Category) (err error) {
	err = checkBlankCategory(data)
	if err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	if _, err := biz.store.GetCategory(ctx, map[string]interface{}{"Categories.CategoryName": data.CategoryName}); err == nil {
		return common.ErrEntityExisted(model.EntityName, common.EntityExisted)
	}

	if err = biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}

func checkBlankCategory(data *model.Category) error {
	data.CategoryName = strings.TrimSpace(data.CategoryName)
	if data.CategoryName == "" {
		return model.ErrCategoryNameIsBlank
	}
	return nil
}
