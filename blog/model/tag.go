package model

import "time"

// Tag struct
type Tag struct {
	id        uint64    `json:"-" gorm:"primaryKey"`
	Posts     []*Post   `json:"posts" gorm:"many2many:posts_params;"`
	Name      string    `validate:"required" json:"name" gorm:"uniqueIndex,type:varchar(80);"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
