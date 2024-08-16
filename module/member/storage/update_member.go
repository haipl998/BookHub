package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateMember(ctx context.Context, data *model.MemberUpdate, cond map[string]interface{}) (err error) {
	var (
		existingMember model.Member
		result         *gorm.DB
	)

	result = s.db.Where("Email = ?", data.Email).Find(&existingMember)
	if result.RowsAffected > 0 {
		return model.ErrEmailExists
	}

	result = s.db.Where("PhoneNumber = ?", data.PhoneNumber).Find(&existingMember)
	if result.RowsAffected > 0 {
		return model.ErrPhoneNumberExists
	}

	if err := s.db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
