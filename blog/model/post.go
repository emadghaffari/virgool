package model

import (
	"time"
)

// StatusPost string for posts status
type StatusPost string

// ScanStatus func
func ScanStatus(value string) StatusPost {
	switch value {
	case "pending":
		return Pending
	case "published":
		return Published
	case "deleted":
		return Deleted
	default:
		return Pending
	}
}

// const for posts status types
const (
	Pending   StatusPost = "pending"
	Published StatusPost = "published"
	Deleted   StatusPost = "deleted"
)

// Post struct
type Post struct {
	ID          uint64     `json:"-" gorm:"primaryKey"`
	UserID      uint64     `validate:"required" json:"user_id"`
	Title       string     `validate:"required" json:"title" gorm:"unique;not null;type:varchar(255);"`
	Slug        string     `validate:"required" json:"slug" gorm:"uniqueIndex;not null;type:varchar(255);"`
	Description string     `validate:"required" json:"description" gorm:"type:text"`
	Text        string     `validate:"required" json:"text" gorm:"type:text"`
	Params      []*Param   `validate:"required" json:"params" gorm:"many2many:params_posts;association_foreignkey:ID;foreignkey:ID"`
	Media       []*Media   `json:"-" gorm:"many2many:posts_medias;association_foreignkey:ID;foreignkey:ID"`
	Tags        []*Tag     `validate:"required" json:"tags" gorm:"many2many:posts_tags;association_foreignkey:ID;foreignkey:ID"`
	Status      StatusPost `validate:"required" json:"status" gorm:"default:pending;not null;type:ENUM('pending', 'published', 'deleted')" sql:"type:ENUM('pending', 'published', 'deleted')"`
	Rate        uint8      `json:"rate" gorm:"gte:1;lte:5;default:1"`
	PublishedAT time.Time  `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
}
