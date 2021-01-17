package model

import "time"

// Query struct
type Query struct {
	Name  string `validate:"required" json:"name" gorm:"not null;type:varchar(80);"`
	Value string `validate:"required" json:"value" gorm:"not null;type:varchar(120);"`
}

// Param struct
type Param struct {
	ID uint64 `json:"-" gorm:"primaryKey"`
	Query
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
