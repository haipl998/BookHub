package storage

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
)

func (s *sqlStore) ListBook(ctx context.Context) (result []model.Book, er error) {
	cond := make(map[string]interface{})
	cond["Books.Deleted"] = false
	cond["Categories.Deleted"] = false
	cond["BookAuthors.Deleted"] = false
	cond["Authors.Deleted"] = false

	db := s.db.Table(model.Book{}.TableName())

	if err := db.Select("Books.BookID, Books.Title, Books.ISBN, Books.PublishedYear, Categories.CategoryName, Authors.FirstName, Authors.LastName").
		Joins("join Categories on Books.CategoryID = Categories.CategoryID").
		Joins("JOIN BookAuthors on Books.BookID = BookAuthors.BookID").
		Joins("join Authors on BookAuthors.AuthorID = Authors.AuthorID").
		Where(cond).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
