package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"errors"

	"context"
)

type DeleteCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (result *model.Category, err error)
	DeleteCategory(ctx context.Context, cond map[string]interface{}) (err error)
}

type deleteCategoryBiz struct {
	store DeleteCategoryStorage
}

func NewDeleteCategoryBiz(store DeleteCategoryStorage) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategoryById(ctx context.Context, id int) (err error) {
	_, err = biz.store.GetCategory(ctx, map[string]interface{}{"Categories.CategoryID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	if err = biz.store.DeleteCategory(ctx, map[string]interface{}{"CategoryID": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
