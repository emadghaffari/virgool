package model

import "time"

// Tag struct
type Tag struct {
	ID        uint64    `json:"-" gorm:"primaryKey"`
	Posts     []*Post   `json:"posts" gorm:"many2many:posts_tags;"`
	Name      string    `validate:"required" json:"name" gorm:"uniqueIndex;type:varchar(40);"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
