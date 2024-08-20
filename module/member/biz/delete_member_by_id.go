package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"
	"errors"

	"context"
)

type DeleteMemberStorage interface {
	GetMember(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error)
	DeleteMemberById(ctx context.Context, cond map[string]interface{}) (err error)
}

type deleteMemberBiz struct {
	store DeleteMemberStorage
}

func NewDeleteMemberBiz(store DeleteMemberStorage) deleteMemberBiz {
	return deleteMemberBiz{store: store}
}

func (biz deleteMemberBiz) DeleteMemberByID(ctx context.Context, id int) (err error) {
	if _, err = biz.store.GetMember(ctx, map[string]interface{}{"Members.MemberID": id}); err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	if err = biz.store.DeleteMemberById(ctx, map[string]interface{}{"Members.MemberID": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
