package storage_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"
	"context"
)

func (s *sqlStore) RegisterMember(ctx context.Context, data *model_member.MemberCreation) (err error) {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
