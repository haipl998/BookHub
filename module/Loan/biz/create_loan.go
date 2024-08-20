package biz

import (
	"BookHub/common"
	"BookHub/module/Loan/model"
	"time"

	"context"
)

type CreateLoanStorage interface {
	CreateLoan(ctx context.Context, data *model.LoanCreation) (err error)
}

type createLoanBiz struct {
	store CreateLoanStorage
}

func NewCreateLoanBiz(store CreateLoanStorage) *createLoanBiz {
	return &createLoanBiz{store: store}
}

func (biz *createLoanBiz) CreateLoan(ctx context.Context, data *model.LoanCreation) (err error) {
	//Todo check ca truong khong duoc phep chong
	if err = checkBlankLoanCreation(data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	// ngay muon la thoi gia hien tai
	data.LoanDate = time.Now()
	// ngay tra la 3 ngay sau do
	data.DueDate = time.Now().Add(time.Duration(72) * time.Hour)

	if err = biz.store.CreateLoan(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}

func checkBlankLoanCreation(data *model.LoanCreation) error {
	if data.BookID == 0 {
		return model.ErrBookIDIsBlank
	}

	if data.MemberID == 0 {
		return model.ErrMemberIDIsBlank
	}

	return nil
}
