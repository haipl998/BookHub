package biz

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
	"errors"
)

type DeleteBookStorage interface {
	GetBookById(ctx context.Context, cond map[string]interface{}) (book *model.Book, err error)
	DeleteBookById(ctx context.Context, id int) (err error)
}

type deleteBookBiz struct {
	store DeleteBookStorage
}

func NewDeleteBookByIdBiz(store DeleteBookStorage) *deleteBookBiz {
	return &deleteBookBiz{store: store}
}

func (biz *deleteBookBiz) DeleteBookById(ctx context.Context, id int) (err error) {
	_, err = biz.store.GetBookById(ctx, map[string]interface{}{"Books.BookID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	if err = biz.store.DeleteBookById(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}
	return nil

}
