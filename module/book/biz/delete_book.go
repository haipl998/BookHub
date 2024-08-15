package biz

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
	"errors"
)

type DeleteBookStorage interface {
	GetBook(ctx context.Context, cond map[string]interface{}) (book *model.Book, err error)
	DeleteBook(ctx context.Context, id int) (err error)
}

type deleteBookBiz struct {
	store DeleteBookStorage
}

func NewDeleteBookByIdBiz(store DeleteBookStorage) *deleteBookBiz {
	return &deleteBookBiz{store: store}
}

func (biz *deleteBookBiz) DeleteBookById(ctx context.Context, id int) (err error) {
	_, err = biz.store.GetBook(ctx, map[string]interface{}{"Books.BookID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrorCannotGetEntity(model.EntityName, err)
		}
		return common.ErrorCannotDeleteEntity(model.EntityName, err)
	}

	if err = biz.store.DeleteBook(ctx, id); err != nil {
		return common.ErrorCannotDeleteEntity(model.EntityName, err)
	}
	return nil

}
