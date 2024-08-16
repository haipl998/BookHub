package model

import (
	"errors"
	"time"
)

const (
	EntityName = "member"
)

var (
	ErrEmailExists       = errors.New("email already exists")
	ErrPhoneNumberExists = errors.New("phone number already exists")
)

type Member struct {
	MemberID    int       `json:"MemberID,omitempty" gorm:"primaryKey;column:MemberID"`
	FirstName   string    `json:"FirstName" gorm:"column:FirstName"`
	LastName    string    `json:"LastName" gorm:"column:LastName"`
	Email       string    `json:"Email" gorm:"column:Email"`
	PhoneNumber string    `json:"PhoneNumber" gorm:"column:PhoneNumber"`
	JoinDate    time.Time `json:"JoinDate" gorm:"column:JoinDate"`
}

type MemberUpdate struct {
	FirstName   string `json:"FirstName,omitempty" gorm:"column:FirstName"`
	LastName    string `json:"LastName,omitempty" gorm:"column:LastName"`
	Email       string `json:"Email,omitempty" gorm:"column:Email"`
	PhoneNumber string `json:"PhoneNumber,omitempty" gorm:"column:PhoneNumber"`
}

func (Member) TableName() string       { return "Members" }
func (MemberUpdate) TableName() string { return Member{}.TableName() }
