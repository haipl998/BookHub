package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"
)

type ListMemberStorage interface {
	ListMember(ctx context.Context) (result *[]model.Member, err error)
}

type listMemberBiz struct {
	store ListMemberStorage
}

func NewListMemberBiz(store ListMemberStorage) *listMemberBiz {
	return &listMemberBiz{store: store}
}

func (biz *listMemberBiz) ListMember(ctx context.Context) (result *[]model.Member, err error) {
	result, err = biz.store.ListMember(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return result, nil

}
