package biz

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
	"errors"
)

type UpdateBookStorage interface {
	GetBookById(ctx context.Context, cond map[string]interface{}) (book *model.Book, err error)
	UpdateBookById(ctx context.Context, cond map[string]interface{}, data *model.BookUpdate) (err error)
}

type updateBookBiz struct {
	store UpdateBookStorage
}

func NewUpdateBookByIdBiz(store UpdateBookStorage) *updateBookBiz {
	return &updateBookBiz{store: store}
}

// TODO: có nên thêm update luôn cả category hoặc tác giả không?
func (biz *updateBookBiz) UpdateBookById(ctx context.Context, id int, data *model.BookUpdate) (err error) {
	_, err = biz.store.GetBookById(ctx, map[string]interface{}{"Books.BookID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if err = biz.store.UpdateBookById(ctx, map[string]interface{}{"Books.BookID": id}, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil

}
