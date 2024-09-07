package biz

import (
	"BookHub/common"
	"BookHub/module/Loan/model"
	"errors"

	"context"
)

type DeleteLoanStorage interface {
	GetLoan(ctx context.Context, cond map[string]interface{}) (result *model.Loan, err error)
	DeleteLoan(ctx context.Context, cond map[string]interface{}) (err error)
}

type deleteLoanBiz struct {
	store DeleteLoanStorage
}

func NewDeleteLoanBiz(store DeleteLoanStorage) *deleteLoanBiz {
	return &deleteLoanBiz{store: store}
}

func (biz *deleteLoanBiz) DeleteLoanById(ctx context.Context, id int) (err error) {
	if _, err = biz.store.GetLoan(ctx, map[string]interface{}{"Loans.LoanID": id}); err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	if err = biz.store.DeleteLoan(ctx, map[string]interface{}{"Loans.LoanID": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
