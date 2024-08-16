package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"errors"
	"strings"

	"context"
)

type UpdateCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (result *model.Category, err error)
	UpdateCategory(ctx context.Context, cond map[string]interface{}, data *model.CategoryUpdate) (err error)
}

type updateCategoryBiz struct {
	store UpdateCategoryStorage
}

func NewUpdateCategoryBiz(store UpdateCategoryStorage) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategoryById(ctx context.Context, id int, data *model.CategoryUpdate) (err error) {
	data.CategoryName = strings.TrimSpace(data.CategoryName)
	if data.CategoryName == "" {
		return common.ErrorCannotUpdaterEntity(model.EntityName, model.ErrCategoryNameIsBlank)
	}

	_, err = biz.store.GetCategory(ctx, map[string]interface{}{"Categories.CategoryID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrorCannotGetEntity(model.EntityName, err)
		}
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}

	if err = biz.store.UpdateCategory(ctx, map[string]interface{}{"Categories.CategoryID": id}, data); err != nil {
		return common.ErrorCannotCreateEntity(model.EntityName, err)
	}

	return nil

}
