package storage_member

import (
	"BookHub/common"
	model_meber "BookHub/module/member/model"

	"context"
)

func (s *sqlStore) ListMember(ctx context.Context) (result *[]model_meber.Member, err error) {
	cond := make(map[string]interface{})
	cond["Members.Deleted"] = false
	db := s.db.Table(model_meber.Member{}.TableName())

	if err := db.Where(cond).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
