package biz_member

import (
	"BookHub/common"
	model_member "BookHub/module/member/model"

	"time"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type RegisterStorage interface {
	GetMember(ctx context.Context, cond map[string]interface{}) (result *model_member.Member, err error)
	RegisterMember(ctx context.Context, data *model_member.MemberCreation) (err error)
}

type registerBiz struct {
	store RegisterStorage
}

func NewRegisterBiz(store RegisterStorage) *registerBiz {
	return &registerBiz{store: store}
}

func (biz *registerBiz) Register(ctx context.Context, data *model_member.MemberCreation) (err error) {
	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"Email": data.Email}); err == nil {
		return common.ErrCannotCreateEntity(model_member.EntityName, model_member.ErrEmailExists)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"PhoneNumber": data.PhoneNumber}); err == nil {
		return common.ErrCannotCreateEntity(model_member.EntityName, model_member.ErrPhoneNumberExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ErrCannotCreateEntity(model_member.EntityName, err)
	}

	data.JoinDate = time.Now()
	data.Password = string(hashedPassword)
	data.Role = "user"

	if err = biz.store.RegisterMember(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model_member.EntityName, err)
	}
	return nil
}
