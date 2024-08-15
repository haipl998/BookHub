package storage

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteBook(ctx context.Context, id int) (err error) {
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("BookAuthors").Where(map[string]interface{}{"BookAuthors.BookID": id}).Delete(nil).Error; err != nil {
			return common.ErrDB(err)
		}

		if err := tx.Table("Loans").Where(map[string]interface{}{"Loans.BookID": id}).Delete(nil).Error; err != nil {
			return common.ErrDB(err)
		}

		if err := tx.Table("Reviews").Where(map[string]interface{}{"Reviews.BookID": id}).Delete(nil).Error; err != nil {
			return common.ErrDB(err)
		}

		if err := tx.Table(model.Book{}.TableName()).Where(map[string]interface{}{"Books.BookID": id}).Delete(nil).Error; err != nil {
			return common.ErrDB(err)
		}

		return nil
	}); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
