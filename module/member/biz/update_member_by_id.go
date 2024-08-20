package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"
	"errors"

	"context"
)

type UpdateMemberStorage interface {
	GetMember(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error)
	UpdateMember(ctx context.Context, data *model.MemberUpdate, cond map[string]interface{}) (err error)
}

type updateMemberBiz struct {
	store UpdateMemberStorage
}

func NewUpdateMemberBiz(store UpdateMemberStorage) *updateMemberBiz {
	return &updateMemberBiz{store: store}
}

func (biz *updateMemberBiz) UpdateMember(ctx context.Context, data *model.MemberUpdate, id int) (err error) {
	_, err = biz.store.GetMember(ctx, map[string]interface{}{"Members.MemberID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"Email": data.Email}); err == nil {
		return common.ErrCannotUpdateEntity(model.EntityName, model.ErrEmailExists)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"PhoneNumber": data.PhoneNumber}); err == nil {
		return common.ErrCannotUpdateEntity(model.EntityName, model.ErrPhoneNumberExists)
	}

	if err = biz.store.UpdateMember(ctx, data, map[string]interface{}{"Members.MemberID": id}); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil

}
