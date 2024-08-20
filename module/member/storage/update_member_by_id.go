package storage

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"
)

func (s *sqlStore) UpdateMember(ctx context.Context, data *model.MemberUpdate, cond map[string]interface{}) (err error) {
	if err := s.db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
