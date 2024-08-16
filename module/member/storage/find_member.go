package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"
	"errors"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetMemberByID(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error) {
	db := s.db.Table(model.Member{}.TableName())

	if err := db.First(&result, cond).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
