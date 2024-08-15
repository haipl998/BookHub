package model

import "errors"

const (
	EntityName = "book"
)

var (
	ErrTitleIsBlank        = errors.New("title cannot be blank")
	ErrCategoryNameIsBlank = errors.New("category name cannot be blank")
	ErrFirstNameIsBlank    = errors.New("first name cannot be blank")
	ErrLastNameIsBlank     = errors.New("last name cannot be blank")
)

type Book struct {
	BookID        int    `json:"BookID,omitempty" gorm:"primaryKey;column:BookID"`
	Title         string `json:"Title" gorm:"column:Title"`
	ISBN          string `json:"ISBN" gorm:"column:ISBN"`
	PublishedYear int    `json:"PublishedYear" gorm:"column:PublishedYear"`
	Categories
	Authors
}

type BookCreation struct {
	BookID        int    `json:"-" gorm:"primaryKey;column:BookID"`
	Title         string `json:"Title" gorm:"column:Title"`
	ISBN          string `json:"ISBN" gorm:"column:ISBN"`
	PublishedYear int    `json:"PublishedYear" gorm:"column:PublishedYear"`
	CategoryID    int    `json:"CategoryID,omitempty" gorm:"column:CategoryID"`
}

type BookUpdate struct {
	Title         string `json:"Title" gorm:"column:Title"`
	ISBN          string `json:"ISBN" gorm:"column:ISBN"`
	PublishedYear int    `json:"PublishedYear" gorm:"column:PublishedYear"`
}

type Categories struct {
	CategoryID   int    `json:"CategoryID,omitempty" gorm:"primaryKey;column:CategoryID"`
	CategoryName string `json:"CategoryName" gorm:"column:CategoryName"`
}

type BookAuthors struct {
	BookID   int `json:"BookID" gorm:"column:BookID"`
	AuthorID int `json:"AuthorID" gorm:"column:AuthorID"`
}

type Authors struct {
	AuthorID  int    `json:"AuthorID,omitempty" gorm:"primaryKey;column:AuthorID"`
	FirstName string `json:"FirstName" gorm:"column:FirstName"`
	LastName  string `json:"LastName" gorm:"column:LastName"`
}

func (Book) TableName() string         { return "Books" }
func (BookCreation) TableName() string { return Book{}.TableName() }
func (BookUpdate) TableName() string   { return Book{}.TableName() }
func (Categories) TableName() string   { return "Categories" }
func (BookAuthors) TableName() string  { return "BookAuthors" }
func (Authors) TableName() string      { return "Authors" }
