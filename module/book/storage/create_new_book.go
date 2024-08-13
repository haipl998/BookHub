package storage

import (
	"BookHub/module/book/model"
	"context"
	"log"

	"gorm.io/gorm"
)

func (s *sqlStore) CreateBook(ctx context.Context, book *model.Book) (err error) {
	// Transaction để thêm sách mới
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Bước 1: Kiểm tra và thêm Category nếu chưa tồn tại
		var category model.Categories
		if err := tx.Table(model.Categories{}.TableName()).Where("CategoryName = ?", book.CategoryName).FirstOrCreate(&category, model.Categories{CategoryName: book.CategoryName}).Error; err != nil {
			log.Printf("Error during FirstOrCreate: %v", err)
			return err
		}

		// Bước 2: Kiểm tra và thêm Author nếu chưa tồn tại
		var author model.Authors
		if err := tx.Table(model.Authors{}.TableName()).Where("FirstName = ? AND LastName = ?", book.FirstName, book.LastName).FirstOrCreate(&author, model.Authors{FirstName: book.FirstName, LastName: book.LastName}).Error; err != nil {
			return err
		}

		// Bước 3: Thêm sách mới vào bảng Books
		bookCreation := model.BookCreation{
			Title:         book.Title,
			ISBN:          book.ISBN,
			PublishedYear: book.PublishedYear,
			CategoryID:    category.CategoryID,
		}
		if err := tx.Table(model.BookCreation{}.TableName()).Create(&bookCreation).Error; err != nil {
			return err
		}
		// lấy bookid để in ra sau thi thêm thành công
		book.BookID = bookCreation.BookID

		// Bước 4: Thêm liên kết giữa sách và tác giả vào bảng BookAuthors
		bookAuthor := model.BookAuthors{
			BookID:   bookCreation.BookID,
			AuthorID: author.AuthorID,
		}
		if err := tx.Table(model.BookAuthors{}.TableName()).Create(&bookAuthor).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
