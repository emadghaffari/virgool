package model

import (
	"time"

	"github.com/emadghaffari/virgool/auth/pkg/grpc/pb"
)

// User struct
type User struct {
	ID        uint64       `json:"id" gorm:"autoIncrementIncrement"`
	Username  string       `json:"username,omitempty"`
	Password  *string      `json:"-,omitempty"`
	Name      string       `json:"name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Phone     string       `json:"phone,omitempty"`
	Email     string       `json:"email,omitempty"`
	Token     string       `json:"token,omitempty"`
	RoleID    uint64       `json:"-"`
	Role      *pb.Role     `json:"roles,omitempty" gorm:"foreignKey:id;references:RoleID"`
	Media     []Mediaables `gorm:"polymorphic:Owner;"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// Role struct
type Role struct {
	ID          uint64        `json:"-" gorm:"autoIncrementIncrement"`
	Name        string        `json:"name,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:roles_permissions;"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// Permission struct
type Permission struct {
	ID        uint64    `json:"-" gorm:"autoIncrementIncrement"`
	Name      string    `json:"name,omitempty"`
	Role      []*Role   `json:"-" gorm:"many2many:roles_permissions;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Media struct
type Media struct {
	ID          uint64    `json:"-" gorm:"autoIncrementIncrement"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	Title       *string   `json:"title,omitempty"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Mediaables struct
type Mediaables struct {
	ID            uint64 `json:"-" gorm:"autoIncrementIncrement"`
	MediaID       uint64 `json:"-"`
	MediaableID   uint64 `json:"-"`
	MediaableType string `json:"-"`
}
