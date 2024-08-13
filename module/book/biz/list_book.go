package biz

import (
	"BookHub/module/book/model"
	"context"
)

type ListBookStorage interface {
	ListBook(ctx context.Context) (result []model.Book, err error)
}

type listBookBiz struct {
	store ListBookStorage
}

func NewListBookBiz(store ListBookStorage) *listBookBiz {
	return &listBookBiz{store: store}
}

func (biz *listBookBiz) ListBook(ctx context.Context) (result []model.Book, err error) {
	result, err = biz.store.ListBook(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil

}
