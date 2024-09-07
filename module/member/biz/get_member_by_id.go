package biz_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"

	"context"
)

type GetMemberStorage interface {
	GetMember(ctx context.Context, cond map[string]interface{}) (result *model_member.Member, err error)
}

type getMemberBiz struct {
	store GetMemberStorage
}

func NewGetMemberBiz(store GetMemberStorage) *getMemberBiz {
	return &getMemberBiz{store: store}
}

func (biz *getMemberBiz) GetMemberByID(ctx context.Context, id int) (result *model_member.Member, err error) {
	result, err = biz.store.GetMember(ctx, map[string]interface{}{"Members.MemberID": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model_member.EntityName, err)
	}
	return result, nil

}
