package model

import "time"

const (
	EntityName = "loan"
)

type Loan struct {
	LoanID     int       `json:"LoanID,omitempty" gorm:"primaryKey;column:LoanID"`
	BookID     int       `json:"BookID" gorm:"column:BookID"`
	MemberID   int       `json:"MemberID" gorm:"column:MemberID"`
	LoanDate   time.Time `json:"LoanDate" gorm:"column:LoanDate"`
	DueDate    time.Time `json:"DueDate" gorm:"column:DueDate"`
	ReturnDate time.Time `json:"ReturnDate" gorm:"column:ReturnDate"`
}

type LoanCreation struct {
	LoanID   int       `json:"LoanID,omitempty" gorm:"primaryKey;column:LoanID"`
	BookID   int       `json:"BookID" gorm:"column:BookID"`
	MemberID int       `json:"MemberID" gorm:"column:MemberID"`
	LoanDate time.Time `json:"LoanDate" gorm:"column:LoanDate"`
	DueDate  time.Time `json:"DueDate" gorm:"column:DueDate"`
}

type LoanUpdate struct {
	DueDate    time.Time `json:"DueDate" gorm:"column:DueDate"`
	ReturnDate time.Time `json:"ReturnDate" gorm:"column:ReturnDate"`
}

func (Loan) TableName() string         { return "Loans" }
func (LoanCreation) TableName() string { return Loan{}.TableName() }
func (LoanUpdate) TableName() string   { return Loan{}.TableName() }
