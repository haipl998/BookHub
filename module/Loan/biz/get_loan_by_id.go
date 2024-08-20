package biz

import (
	"BookHub/common"
	"BookHub/module/Loan/model"

	"context"
)

type GetLoanStorage interface {
	GetLoan(ctx context.Context, cond map[string]interface{}) (result *model.Loan, err error)
}

type getLoanBiz struct {
	store GetLoanStorage
}

func NewGetLoanBiz(store GetLoanStorage) *getLoanBiz {
	return &getLoanBiz{store: store}
}

func (biz *getLoanBiz) GetLoan(ctx context.Context, id int) (result *model.Loan, err error) {
	result, err = biz.store.GetLoan(ctx, map[string]interface{}{"Loans.LoanID": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}
	return result, nil

}
