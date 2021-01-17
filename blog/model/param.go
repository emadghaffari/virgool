package model

import "time"

// Query struct
type Query struct {
	Name  string `validate:"required" json:"name" gorm:"type:varchar(80);"`
	Value string `validate:"required" json:"value" gorm:"type:varchar(120);"`
}

// Param struct
type Param struct {
	Query
	ID        uint64    `json:"-" gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
