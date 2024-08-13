package storage

import (
	"BookHub/module/book/model"
	"context"
)

func (s *sqlStore) ListBook(ctx context.Context) (result []model.Book, er error) {
	db := s.db.Table(model.Book{}.TableName())

	if err := db.Select("Books.BookID, Books.Title, Books.ISBN, Books.PublishedYear, Categories.CategoryName, Authors.FirstName, Authors.LastName").
		Joins("join Categories on Books.CategoryID = Categories.CategoryID").Joins("JOIN BookAuthors on Books.BookID = BookAuthors.BookID").
		Joins("join Authors on BookAuthors.AuthorID = Authors.AuthorID").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
