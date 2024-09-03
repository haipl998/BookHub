package storage_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"

	"context"
)

func (s *sqlStore) UpdateMember(ctx context.Context, data *model_member.MemberUpdate, cond map[string]interface{}) (err error) {
	if err := s.db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
