package storage

import (
	"BookHub/common"
	"BookHub/module/Loan/model"

	"context"
)

func (s *sqlStore) ListLoan(ctx context.Context) (result *[]model.Loan, err error) {
	db := s.db.Table(model.Loan{}.TableName())

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
