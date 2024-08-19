package storage

import (
	"BookHub/common"
	"BookHub/module/Loan/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) GetLoan(ctx context.Context, cond map[string]interface{}) (result *model.Loan, err error) {
	if err := s.db.First(&result, cond).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
