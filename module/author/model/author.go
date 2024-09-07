package model

import (
	"errors"
	"strings"
)

const (
	EntityName = "Author"
)

var (
	ErrFirstNameIsBlank = errors.New("first name cannot be blank")
	ErrLastNameIsBlank  = errors.New("last name cannot be blank")
	ErrBothIsBlank      = errors.New("both first name and last name be blank")
)

type Author struct {
	AuthorID  int    `json:"AuthorID,omitempty" gorm:"primaryKey;column:AuthorID"`
	FirstName string `json:"FirstName" gorm:"column:FirstName"`
	LastName  string `json:"LastName" gorm:"column:LastName"`
}

type AuthorUpdate struct {
	FirstName string `json:"FirstName,omitempty" gorm:"column:FirstName"`
	LastName  string `json:"LastName,omitempty" gorm:"column:LastName"`
}

func (Author) TableName() string       { return "Authors" }
func (AuthorUpdate) TableName() string { return Author{}.TableName() }

func (a *Author) Validate() error {
	if strings.TrimSpace(a.FirstName) == "" {
		return ErrFirstNameIsBlank
	}

	if strings.TrimSpace(a.LastName) == "" {
		return ErrLastNameIsBlank
	}

	return nil
}

func (au *AuthorUpdate) Validate() error {
	if au.FirstName != "" && strings.TrimSpace(au.FirstName) == "" {
		return ErrFirstNameIsBlank
	}

	if au.LastName != "" && strings.TrimSpace(au.LastName) == "" {
		return ErrLastNameIsBlank
	}

	return nil
}
