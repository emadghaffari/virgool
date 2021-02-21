package model

import "time"

// Point struct
type Point struct {
	ID        uint64    `json:"-" gorm:"primaryKey"`
	UserID    uint64    `validate:"required" json:"user_id"  gorm:"unique;not null"`
	Point     uint64    `validate:"required" json:"point"`
	UpdatedAt time.Time `json:"-"`
}

// Details struct
type Details struct {
	ID          uint64    `json:"-" gorm:"primaryKey"`
	PointID     uint64    `json:"-"`
	point       Point     `gorm:"foreignKey:PointID;references:ID"`
	Point       int       `json:"point"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
