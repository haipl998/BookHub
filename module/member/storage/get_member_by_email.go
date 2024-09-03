package storage_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *sqlStore) GetMemberByEmail(ctx context.Context, cond map[string]interface{}) (result *model_member.SessionMember, err error) {
	cond["Members.Deleted"] = false

	db := s.db.Table(model_member.SessionMember{}.TableName())

	if err := db.Where(cond).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
