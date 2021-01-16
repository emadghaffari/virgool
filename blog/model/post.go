package model

import (
	"database/sql/driver"
	"time"
)

// StatusPost string for posts status
type StatusPost string

// Scan the post status
func (e *StatusPost) Scan(value interface{}) error {
	*e = StatusPost(value.([]byte))
	return nil
}

// Value return the post status
func (e StatusPost) Value() (driver.Value, error) {
	return string(e), nil
}

// const for posts status types
const (
	Pending   StatusPost = "pending"
	Published StatusPost = "published"
	Deleted   StatusPost = "deleted"
)

// Post struct
type Post struct {
	id          uint64     `json:"-" gorm:"primaryKey"`
	UserID      uint64     `validate:"required" json:"user_id"`
	Title       string     `validate:"required" json:"title" gorm:"unique;not null;type:varchar(180);"`
	Slug        string     `validate:"required" json:"slug" gorm:"uniqueIndex;not null;type:varchar(250);"`
	Description string     `validate:"required" json:"description"`
	Text        string     `validate:"required" json:"text"`
	Params      []*Param   `validate:"required" json:"params" gorm:"many2many:posts_params;association_foreignkey:ID;foreignkey:ID"`
	Media       []*Media   `json:"-" gorm:"many2many:posts_medias;association_foreignkey:ID;foreignkey:ID"`
	Tags        []*Tag     `validate:"required" json:"tags" gorm:"many2many:posts_tags;association_foreignkey:ID;foreignkey:ID"`
	Status      StatusPost `validate:"required" json:"status" gorm:"default:pending" sql:"type:ENUM('pending', 'published', 'deleted')"`
	Rate        uint8      `json:"rate" gorm:"gte:1;lte:5;default:1"`
	PublishedAT time.Time  `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
}
