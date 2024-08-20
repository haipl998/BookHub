package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"
)

func (s *sqlStore) CreateMember(ctx context.Context, data *model.Member) (err error) {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
