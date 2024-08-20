package biz

import (
	"BookHub/common"
	"BookHub/module/member/model"
	"strings"
	"time"

	"context"
)

type CreateMemberStorage interface {
	GetMember(ctx context.Context, cond map[string]interface{}) (result *model.Member, err error)
	CreateMember(ctx context.Context, data *model.Member) (err error)
}

type createMemberBiz struct {
	store CreateMemberStorage
}

func NewCreateMemberBiz(store CreateMemberStorage) *createMemberBiz {
	return &createMemberBiz{store: store}
}

func (biz *createMemberBiz) CreateMember(ctx context.Context, data *model.Member) (err error) {
	if err := checkBlankMember(data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"Email": data.Email}); err == nil {
		return common.ErrCannotCreateEntity(model.EntityName, model.ErrEmailExists)
	}

	if _, err := biz.store.GetMember(ctx, map[string]interface{}{"PhoneNumber": data.PhoneNumber}); err == nil {
		return common.ErrCannotCreateEntity(model.EntityName, model.ErrPhoneNumberExists)
	}

	data.JoinDate = time.Now()

	if err = biz.store.CreateMember(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}

func checkBlankMember(data *model.Member) error {
	data.Email = strings.TrimSpace(data.Email)
	if data.Email == "" {
		return model.ErrEmailIsBlank
	}

	data.PhoneNumber = strings.TrimSpace(data.PhoneNumber)
	if data.PhoneNumber == "" {
		return model.ErrPhoneNumberIsBlank
	}

	data.FirstName = strings.TrimSpace(data.FirstName)
	if data.FirstName == "" {
		return model.ErrFirstNameIsBlank
	}

	data.LastName = strings.TrimSpace(data.LastName)
	if data.LastName == "" {
		return model.ErrLastNameIsBlank
	}

	return nil
}
