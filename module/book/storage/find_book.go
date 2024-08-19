package storage

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) GetBookById(ctx context.Context, cond map[string]interface{}) (book *model.Book, err error) {
	cond["Books.Deleted"] = false
	cond["Categories.Deleted"] = false
	cond["BookAuthors.Deleted"] = false
	cond["Authors.Deleted"] = false
	db := s.db.Table(model.Book{}.TableName())

	if err := db.Select("Books.BookID, Books.Title, Books.ISBN, Books.PublishedYear, Categories.CategoryName, Authors.FirstName, Authors.LastName").
		Joins("join Categories on Books.CategoryID = Categories.CategoryID").Joins("JOIN BookAuthors on Books.BookID = BookAuthors.BookID").
		Joins("join Authors on BookAuthors.AuthorID = Authors.AuthorID").
		Where(cond).
		First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return book, nil
}
