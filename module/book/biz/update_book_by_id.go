package biz

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
	"errors"
)

type UpdateBookStorage interface {
	GetBook(ctx context.Context, cond map[string]interface{}) (book *model.Book, err error)
	UpdateBook(ctx context.Context, cond map[string]interface{}, data *model.BookUpdate) (err error)
}

type updateBookBiz struct {
	store UpdateBookStorage
}

func NewUpdateBookByIdBiz(store UpdateBookStorage) *updateBookBiz {
	return &updateBookBiz{store: store}
}

func (biz *updateBookBiz) UpdateBookById(ctx context.Context, id int, data *model.BookUpdate) (err error) {
	_, err = biz.store.GetBook(ctx, map[string]interface{}{"Books.BookID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrorCannotGetEntity(model.EntityName, err)
		}
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}

	if err = biz.store.UpdateBook(ctx, map[string]interface{}{"Books.BookID": id}, data); err != nil {
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}
	return nil

}
