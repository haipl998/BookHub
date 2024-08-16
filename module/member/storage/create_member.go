package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) CreateMember(ctx context.Context, data *model.Member) (err error) {
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

	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
