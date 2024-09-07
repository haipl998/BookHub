package biz_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"

	"context"
)

type ListMemberStorage interface {
	ListMember(ctx context.Context) (result *[]model_member.Member, err error)
}

type listMemberBiz struct {
	store ListMemberStorage
}

func NewListMemberBiz(store ListMemberStorage) *listMemberBiz {
	return &listMemberBiz{store: store}
}

func (biz *listMemberBiz) ListMember(ctx context.Context) (result *[]model_member.Member, err error) {
	result, err = biz.store.ListMember(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(model_member.EntityName, err)
	}
	return result, nil

}
