package model_member

import (
	"errors"
	"log"
	"strings"
	"time"
)

const (
	EntityName = "member"
)

var (
	ErrEmailExists       = errors.New("email already exists")
	ErrPhoneNumberExists = errors.New("phone number already exists")
	ErrFirstNameBlank    = errors.New("first name cannot be blank")
	ErrLastNameBlank     = errors.New("last name cannot be blank")
	ErrEmailBlank        = errors.New("email cannot be blank")
	ErrPhoneNumberBlank  = errors.New("phone number cannot be blank")
	ErrPasswordBlank     = errors.New("password cannot be blank")
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
	Password    string `json:"Password"`
}

type MemberCreation struct {
	MemberID    int       `json:"MemberID,omitempty" gorm:"primaryKey;column:MemberID"`
	FirstName   string    `json:"FirstName" gorm:"column:FirstName"`
	LastName    string    `json:"LastName" gorm:"column:LastName"`
	Email       string    `json:"Email" gorm:"column:Email"`
	PhoneNumber string    `json:"PhoneNumber" gorm:"column:PhoneNumber"`
	JoinDate    time.Time `json:"JoinDate" gorm:"column:JoinDate"`
	Password    string    `json:"Password"`
	Role        string    `json:"Role"`
}

type SessionMember struct {
	MemberID    int    `json:"MemberID,omitempty" gorm:"primaryKey;column:MemberID"`
	Email       string `json:"Email" gorm:"column:Email"`
	PhoneNumber string `json:"PhoneNumber" gorm:"column:PhoneNumber"`
	Password    string `json:"-"`
	Role        string `json:"Role"`
}

type LoginForm struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (Member) TableName() string         { return "Members" }
func (MemberUpdate) TableName() string   { return Member{}.TableName() }
func (MemberCreation) TableName() string { return Member{}.TableName() }
func (SessionMember) TableName() string  { return Member{}.TableName() }

func (mc *MemberCreation) Validate() error {
	log.Print(mc)
	if strings.TrimSpace(mc.FirstName) == "" {
		return ErrFirstNameBlank
	}
	if strings.TrimSpace(mc.LastName) == "" {
		return ErrLastNameBlank
	}
	if mc.Email == "" {
		return ErrEmailBlank
	}
	if mc.PhoneNumber == "" {
		return ErrPhoneNumberBlank
	}
	if mc.Password == "" {
		return ErrPasswordBlank
	}

	// Add more validations as needed, e.g.:
	// - Password strength check

	return nil
}

func (mc *MemberUpdate) Validate() error {
	if mc.FirstName != "" && strings.TrimSpace(mc.FirstName) == "" {
		return ErrFirstNameBlank
	}
	if mc.LastName != "" && strings.TrimSpace(mc.LastName) == "" {
		return ErrLastNameBlank
	}
	if mc.Email != "" && strings.TrimSpace(mc.Email) == "" {
		return ErrEmailBlank
	}
	if mc.PhoneNumber != "" && strings.TrimSpace(mc.PhoneNumber) == "" {
		return ErrPhoneNumberBlank
	}
	// Kiểm tra mật khẩu nếu được cung cấp
	if mc.Password != "" && strings.TrimSpace(mc.Password) == "" {
		return ErrPasswordBlank
	}
	// Thêm các kiểm tra khác nếu cần, ví dụ:
	// - Kiểm tra độ mạnh của mật khẩu

	return nil
}
