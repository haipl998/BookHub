package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"
	"errors"

	"context"
)

type UpdateMemberStorage interface {
	GetMemberByID(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error)
	UpdateMember(ctx context.Context, data *model.MemberUpdate, cond map[string]interface{}) (err error)
}

type updateMemberBiz struct {
	store UpdateMemberStorage
}

func NewUpdateMemberBiz(store UpdateMemberStorage) *updateMemberBiz {
	return &updateMemberBiz{store: store}
}

func (biz *updateMemberBiz) UpdateMember(ctx context.Context, data *model.MemberUpdate, id int) (err error) {
	//Todo: check is blank

	_, err = biz.store.GetMemberByID(ctx, map[string]interface{}{"Members.MemberID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrorCannotGetEntity(model.EntityName, err)
		}
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}

	if err = biz.store.UpdateMember(ctx, data, map[string]interface{}{"Members.MemberID": id}); err != nil {
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}
	return nil

}
