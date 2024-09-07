package model

import (
	"errors"
	"strings"
)

const (
	EntityName = "category"
)

var (
	ErrCategoryNameIsBlank = errors.New("category name cannot be blank")
)

type Category struct {
	CategoryID   int    `json:"CategoryID,omitempty" gorm:"primaryKey;column:CategoryID"`
	CategoryName string `json:"CategoryName" gorm:"column:CategoryName"`
}

type CategoryUpdate struct {
	CategoryName string `json:"CategoryName" gorm:"column:CategoryName"`
}

func (Category) TableName() string       { return "Categories" }
func (CategoryUpdate) TableName() string { return Category{}.TableName() }

func (c *Category) Validate() error {
	if strings.TrimSpace(c.CategoryName) == "" {
		return ErrCategoryNameIsBlank
	}
	return nil
}

func (cu *CategoryUpdate) Validate() error {
	if strings.TrimSpace(cu.CategoryName) == "" {
		return ErrCategoryNameIsBlank
	}
	return nil
}
