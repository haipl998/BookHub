package storage

import (
	"BookHub/module/book/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteBook(ctx context.Context, id int) (err error) {
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table(model.BookAuthors{}.TableName()).Where(map[string]interface{}{"BookAuthors.BookID": id}).Delete(nil).Error; err != nil {
			return err
		}

		if err := tx.Table(model.Book{}.TableName()).Where(map[string]interface{}{"Books.BookID": id}).Delete(nil).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
