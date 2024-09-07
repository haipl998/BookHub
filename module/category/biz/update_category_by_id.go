package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"errors"

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
	_, err = biz.store.GetCategory(ctx, map[string]interface{}{"Categories.CategoryID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	if _, err := biz.store.GetCategory(ctx, map[string]interface{}{"Categories.CategoryName": data.CategoryName}); err == nil {
		return common.ErrEntityExisted(model.EntityName, common.EntityExisted)
	}

	if err = biz.store.UpdateCategory(ctx, map[string]interface{}{"Categories.CategoryID": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil

}
