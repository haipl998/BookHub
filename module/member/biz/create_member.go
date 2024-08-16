package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"

	"context"
)

type CreateMemberStorage interface {
	CreateMember(ctx context.Context, data *model.Member) (err error)
}

type createMemberBiz struct {
	store CreateMemberStorage
}

func NewCreateMemberBiz(store CreateMemberStorage) *createMemberBiz {
	return &createMemberBiz{store: store}
}

func (biz *createMemberBiz) CreateMember(ctx context.Context, data *model.Member) (err error) {
	//Todo: check is blank
	if err = biz.store.CreateMember(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(model.EntityName, err)
	}
	return nil

}
