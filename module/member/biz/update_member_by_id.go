package biz_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"
	"errors"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type UpdateMemberStorage interface {
	GetMember(ctx context.Context, cond map[string]interface{}) (result *model_member.Member, err error)
	UpdateMember(ctx context.Context, data *model_member.MemberUpdate, cond map[string]interface{}) (err error)
}

type updateMemberBiz struct {
	store UpdateMemberStorage
}

func NewUpdateMemberBiz(store UpdateMemberStorage) *updateMemberBiz {
	return &updateMemberBiz{store: store}
}

func (biz *updateMemberBiz) UpdateMember(ctx context.Context, data *model_member.MemberUpdate, id int) (err error) {
	_, err = biz.store.GetMember(ctx, map[string]interface{}{"Members.MemberID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model_member.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model_member.EntityName, err)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"Email": data.Email}); err == nil {
		return common.ErrCannotUpdateEntity(model_member.EntityName, model_member.ErrEmailExists)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"PhoneNumber": data.PhoneNumber}); err == nil {
		return common.ErrCannotUpdateEntity(model_member.EntityName, model_member.ErrPhoneNumberExists)
	}

	if data.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return common.ErrCannotUpdateEntity(model_member.EntityName, err)
		}
		data.Password = string(hashedPassword)
	}

	if err = biz.store.UpdateMember(ctx, data, map[string]interface{}{"Members.MemberID": id}); err != nil {
		return common.ErrCannotUpdateEntity(model_member.EntityName, err)
	}
	return nil

}
