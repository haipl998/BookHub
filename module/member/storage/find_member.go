package storage_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"
	"errors"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetMember(ctx context.Context, cond map[string]interface{}) (result *model_member.Member, err error) {
	cond["Members.Deleted"] = false

	db := s.db.Table(model_member.Member{}.TableName())

	if err := db.Where(cond).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
