package biz

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"

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
	// check xem co ton tai theo id khong
	currentAuthor, err := biz.store.GetAuthor(ctx, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if err := data.Validate(); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	// Tạo điều kiện kiểm tra dựa trên dữ liệu cập nhật và dữ liệu hiện tại
	checkConditionExisted := make(map[string]interface{})
	if data.FirstName != "" {
		checkConditionExisted["Authors.FirstName"] = data.FirstName
	} else {
		checkConditionExisted["Authors.FirstName"] = currentAuthor.FirstName
	}
	if data.LastName != "" {
		checkConditionExisted["Authors.LastName"] = data.LastName
	} else {
		checkConditionExisted["Authors.LastName"] = currentAuthor.LastName
	}

	//check xem ten da co chua
	if _, err := biz.store.GetAuthor(ctx, checkConditionExisted); err == nil {
		return common.ErrEntityExisted(model.EntityName, common.EntityExisted)
	}

	err = biz.store.UpdateAuthorById(ctx, data, map[string]interface{}{"Authors.AuthorID": id})
	if err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}
