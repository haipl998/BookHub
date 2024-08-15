package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

type GetAuthorStorage interface {
	GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Authors, err error)
}

type getAuthorBiz struct {
	store GetAuthorStorage
}

func NewGetAuthorBiz(store GetAuthorStorage) *getAuthorBiz {
	return &getAuthorBiz{store: store}
}

func (biz *getAuthorBiz) GetAuthorById(ctx context.Context, id int) (result *model.Authors, err error) {
	result, err = biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		return nil, common.ErrorCannotGetEntity(model.EntityName, err)
	}
	return result, nil

}
