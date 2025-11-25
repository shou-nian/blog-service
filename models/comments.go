package models

import "gorm.io/gorm"

type Comments struct {
	gorm.Model

	Content       string    `json:"content" gorm:"not null;size:255"`
	ArticlesId    uint      `json:"articles_id" gorm:"not null"`
	Article       Articles  `json:"article" gorm:"foreignKey:ArticlesId;references:ID"`
	UserId        uint      `json:"user_id" gorm:"not null;index"`
	User          Users     `json:"user" gorm:"foreignKey:UserId;references:ID"`
	ParentId      uint      `json:"parent_id" gorm:"index"` // 允许为null，表示顶级评论
	ParentComment *Comments `json:"parent_comment" gorm:"foreignKey:ParentId;references:ID"`
}

func (Comments) TableName() string { return "comments" }
