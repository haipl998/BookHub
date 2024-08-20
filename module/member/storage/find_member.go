package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"
	"errors"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetMember(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error) {
	cond["Members.Deleted"] = false

	db := s.db.Table(model.Member{}.TableName())

	if err := db.Where(cond).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
