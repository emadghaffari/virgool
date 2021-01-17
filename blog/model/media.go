package model

import "time"

// Media struct
type Media struct {
	ID          uint64    `json:"-" gorm:"primaryKey"`
	URL         string    `validate:"required" json:"url" gorm:"unique;not null;type:varchar(500);"`
	Type        string    `validate:"required" json:"type" gorm:"type:varchar(50);"`
	Title       *string   `validate:"required" json:"title" gorm:"type:varchar(255);"`
	Description *string   `validate:"required" json:"description"`
	Post        []*Post   `json:"posts" gorm:"many2many:posts_medias;association_foreignkey:ID;foreignkey:ID"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
