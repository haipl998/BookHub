package storage

import (
	"BookHub/common"
	"BookHub/module/Loan/model"

	"context"
)

func (s *sqlStore) CreateLoan(ctx context.Context, data *model.LoanCreation) (err error) {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
