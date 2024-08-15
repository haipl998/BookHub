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

	if err = biz.store.CreateBook(ctx, book); err != nil {
		return common.ErrorCannotCreateEntity(model.EntityName, err)
	}
	return nil

}
