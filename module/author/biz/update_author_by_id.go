package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"
	"strings"

	"context"
)

type UpdateAuthorStorage interface {
	GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Author, err error)
	UpdateAuthorById(ctx context.Context, data *model.AuthorUpdate, cond map[string]interface{}) (err error)
}

type updateAuthorBiz struct {
	store UpdateAuthorStorage
}

func NewUpdateAuthorBiz(store UpdateAuthorStorage) *updateAuthorBiz {
	return &updateAuthorBiz{store: store}
}

func (biz *updateAuthorBiz) UpdateAuthorById(ctx context.Context, data *model.AuthorUpdate, id int) (err error) {
	err = checkBlankAuthorUpdate(data)
	if err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	// check xem co ton tai theo id khong
	_, err = biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	//check xem ten da co chua
	if _, err := biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.FirstName": data.FirstName, "Authors.LastName": data.LastName}); err == nil {
		return common.ErrEntityExisted(model.EntityName, common.EntityExisted)
	}

	err = biz.store.UpdateAuthorById(ctx, data, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}

// TODO: tìm các viết bằng method
func checkBlankAuthorUpdate(data *model.AuthorUpdate) error {
	data.FirstName = strings.TrimSpace(data.FirstName)
	data.LastName = strings.TrimSpace(data.LastName)
	if data.LastName == "" && data.FirstName == "" {
		return model.ErrBothIsBlank
	}

	return nil
}
