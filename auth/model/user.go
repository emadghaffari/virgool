package model

import (
	"time"
)

// User struct
type User struct {
	ID        uint64       `json:"id" gorm:"primaryKey"`
	Username  string       `validate:"required" json:"username,omitempty" gorm:"unique;not null;type:varchar(100);"`
	Password  *string      `validate:"required,gte=7" json:"-" gorm:"type:varchar(100);"`
	Name      string       `validate:"required" json:"name,omitempty" gorm:"type:varchar(100);"`
	LastName  string       `validate:"required" json:"last_name,omitempty" gorm:"type:varchar(100);"`
	Phone     string       `validate:"required" json:"phone,omitempty" gorm:"unique;not null;type:varchar(30);"`
	Email     string       `validate:"required,email" json:"email,omitempty" gorm:"unique;not null;type:varchar(60);"`
	Token     string       `json:"-"`
	RoleID    uint64       `json:"-"`
	Role      Role         `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	Media     []Mediaables `json:"-" gorm:"polymorphic:Owner;"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"-"`
}

// Role struct
type Role struct {
	ID          uint64        `json:"-" gorm:"primaryKey"`
	Name        string        `json:"name" gorm:"unique;not null;type:varchar(30);"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:roles_permissions;association_foreignkey:ID;foreignkey:ID"`
	CreatedAt   time.Time     `json:"-"`
	UpdatedAt   time.Time     `json:"-"`
}

// Permission struct
type Permission struct {
	ID        uint64    `json:"-" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null;type:varchar(30);"`
	Role      []*Role   `json:"-" gorm:"many2many:roles_permissions;association_foreignkey:ID;foreignkey:ID"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Media struct
type Media struct {
	ID          uint64    `json:"-" gorm:"primaryKey"`
	URL         string    `validate:"required" json:"url"`
	Type        string    `validate:"required" json:"type"`
	Title       *string   `validate:"required" json:"title"`
	Description *string   `validate:"required" json:"description"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

// Mediaables struct
type Mediaables struct {
	ID        uint64 `json:"-" gorm:"primaryKey"`
	MediaID   uint64 `json:"-"`
	OwnerID   uint64 `json:"-"`
	OwnerType string `json:"-"`
}
