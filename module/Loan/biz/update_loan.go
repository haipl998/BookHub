package biz

import (
	"BookHub/common"
	"BookHub/module/Loan/model"
	"errors"
	"time"

	"context"
)

type UpdateLoanStorage interface {
	GetLoan(ctx context.Context, cond map[string]interface{}) (result *model.Loan, err error)
	UpdateLoan(ctx context.Context, data *model.LoanUpdate, cond map[string]interface{}) (err error)
}

type updateLoanBiz struct {
	store UpdateLoanStorage
}

func NewUpdateLoanBiz(store UpdateLoanStorage) *updateLoanBiz {
	return &updateLoanBiz{store: store}
}

func (biz *updateLoanBiz) UpdateLoanById(ctx context.Context, data *model.LoanUpdate, id int) (err error) {
	_, err = biz.store.GetLoan(ctx, map[string]interface{}{"Loans.LoanID": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.ErrorCannotGetEntity(model.EntityName, err)
		}
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}

	// Todo: duedate gia han them 3 ngay
	// tra thi lay thoi gia hien tai
	data.ReturnDate = time.Now()

	if err = biz.store.UpdateLoan(ctx, data, map[string]interface{}{"Loans.LoanID": id}); err != nil {
		return common.ErrorCannotUpdaterEntity(model.EntityName, err)
	}
	return nil

}
