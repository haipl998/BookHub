package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"
)

type GetMemberStorage interface {
	GetMemberByID(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error)
}

type getMemberBiz struct {
	store GetMemberStorage
}

func NewGetMemberBiz(store GetMemberStorage) *getMemberBiz {
	return &getMemberBiz{store: store}
}

func (biz *getMemberBiz) GetMemberByID(ctx context.Context, id int) (result *model.Member, err error) {
	result, err = biz.store.GetMemberByID(ctx, map[string]interface{}{"Members.MemberID": id})
	if err != nil {
		return nil, common.ErrorCannotGetEntity(model.EntityName, err)
	}
	return result, nil

}
