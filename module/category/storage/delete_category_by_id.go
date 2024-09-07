package storage

import (
	"BookHub/common"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, cond map[string]interface{}) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		tables := []string{"Categories", "Books", "BookAuthors", "Loans", "Reviews"}
		for _, table := range tables {
			if err := softDeleteFromTable(tx, table, cond); err != nil {
				return err
			}
		}
		return nil
	})
}

func softDeleteFromTable(tx *gorm.DB, table string, cond map[string]interface{}) error {
	query := tx.Table(table)
	if table != "Categories" && table != "Books" {
		query = query.Where("BookID IN (?)", tx.Table("Books").Select("BookID").Where(cond))
	} else {
		query = query.Where(cond)
	}
	if err := query.Update("Deleted", true).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
