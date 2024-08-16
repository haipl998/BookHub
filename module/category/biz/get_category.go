package biz

import (
	"BookHub/common"
	"BookHub/module/category/model"

	"context"
)

type GetCategoryStorage interface {
	GetCategory(ctx context.Context, cond map[string]interface{}) (result *model.Category, err error)
}

type getCategoryBiz struct {
	store GetCategoryStorage
}

func NewGetCategoryBiz(store GetCategoryStorage) *getCategoryBiz {
	return &getCategoryBiz{store: store}
}

func (biz *getCategoryBiz) GetCategoryById(ctx context.Context, id int) (result *model.Category, err error) {
	result, err = biz.store.GetCategory(ctx, map[string]interface{}{"Categories.CategoryID": id})
	if err != nil {
		return nil, common.ErrorCannotGetEntity(model.EntityName, err)
	}
	return result, nil

}
