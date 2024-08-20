package biz

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
	"strings"
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
	err = checkBlank(book)
	if err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	if err = biz.store.CreateBook(ctx, book); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil

}

func checkBlank(book *model.Book) error {
	book.Title = strings.TrimSpace(book.Title)
	if book.Title == "" {
		return model.ErrTitleIsBlank
	}

	book.CategoryName = strings.TrimSpace(book.CategoryName)
	if book.CategoryName == "" {
		return model.ErrCategoryNameIsBlank
	}

	book.FirstName = strings.TrimSpace(book.FirstName)
	if book.FirstName == "" {
		return model.ErrFirstNameIsBlank
	}

	book.LastName = strings.TrimSpace(book.LastName)
	if book.LastName == "" {
		return model.ErrLastNameIsBlank
	}

	return nil
}
