package model

import "errors"

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
