package biz

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
)

type GetBookStorage interface {
	GetBook(ctx context.Context, cond map[string]interface{}) (book *model.Book, err error)
}

type getBookBiz struct {
	store GetBookStorage
}

func NewGetBookByIdBiz(store GetBookStorage) *getBookBiz {
	return &getBookBiz{store: store}
}

func (biz *getBookBiz) GetBookById(ctx context.Context, id int) (book *model.Book, err error) {
	book, err = biz.store.GetBook(ctx, map[string]interface{}{"Books.BookID": id})
	if err != nil {
		return nil, common.ErrorCannotGetEntity(model.EntityName, err)
	}
	return book, nil

}
