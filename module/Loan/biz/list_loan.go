package biz

import (
	"BookHub/common"
	"BookHub/module/Loan/model"

	"context"
)

type ListLoanStorage interface {
	ListLoan(ctx context.Context) (result *[]model.Loan, err error)
}

type listLoanBiz struct {
	store ListLoanStorage
}

func NewListLoanBiz(store ListLoanStorage) *listLoanBiz {
	return &listLoanBiz{store: store}
}

func (biz *listLoanBiz) ListLoan(ctx context.Context) (result *[]model.Loan, err error) {
	result, err = biz.store.ListLoan(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}
	return result, nil

}
