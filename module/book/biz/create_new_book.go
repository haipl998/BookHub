package biz

import (
	"BookHub/module/book/model"
	"context"
)

type CreateBookStorage interface {
	CreateBook(ctx context.Context, book *model.Book) (err error)
}

type createBookBiz struct {
	store CreateBookStorage
}

func NewCreateBookBiz(store CreateBookStorage) *createBookBiz {
	return &createBookBiz{store: store}
}

func (biz *createBookBiz) CreateBook(ctx context.Context, book *model.Book) (err error) {
	if err = biz.store.CreateBook(ctx, book); err != nil {
		return err
	}
	return nil

}
