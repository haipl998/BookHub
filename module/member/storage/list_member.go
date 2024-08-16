package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"
)

func (s *sqlStore) ListMember(ctx context.Context) (result *[]model.Member, err error) {
	db := s.db.Table(model.Member{}.TableName())

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
