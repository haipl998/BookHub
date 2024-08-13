package model

type Book struct {
	BookID        int    `json:"BookID" gorm:"column:BookID"`
	Title         string `json:"Title" gorm:"column:Title"`
	ISBN          string `json:"ISBN" gorm:"column:ISBN"`
	PublishedYear int    `json:"PublishedYear" gorm:"column:PublishedYear"`
	Categorys
	Authors
}

type Categorys struct {
	CategoryID   int    `json:"CategoryID,omitempty" gorm:"column:CategoryID"`
	CategoryName string `json:"CategoryName" gorm:"column:CategoryName"`
}

type Authors struct {
	AuthorID  int    `json:"AuthorID,omitempty" gorm:"column:AuthorID"`
	FirstName string `json:"FirstName" gorm:"column:FirstName"`
	LastName  string `json:"LastName" gorm:"column:LastName"`
}

func (Book) TableName() string { return "Books" }
