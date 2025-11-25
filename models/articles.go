package models

import "gorm.io/gorm"

type status string

const (
	Published status = "published"
	Draft     status = "draft"
)

type Articles struct {
	gorm.Model

	Title    string `json:"title" gorm:"not null;size:255"`
	Content  string `json:"content" gorm:"not null"`
	Summary  string `json:"summary" gorm:"size:255"`
	AuthorID uint   `json:"author_id" gorm:"not null;index"`
	Author   Users  `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	Status   status `json:"status" gorm:"default:draft"`
}

func (Articles) TableName() string { return "articles" }
