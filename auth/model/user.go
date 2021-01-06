package model

import (
	"time"

	"github.com/emadghaffari/virgool/auth/pkg/grpc/pb"
)

// User struct
type User struct {
	ID        uint64    `json:"-" sql:"index"`
	Username  string    `json:"username,omitempty"`
	Password  *string   `json:"-,omitempty"`
	Name      string    `json:"name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Email     string    `json:"email,omitempty"`
	Token     *string   `json:"token,omitempty"`
	Role      *pb.Role  `json:"roles,omitempty" gorm:"foreignKey:id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Role struct
type Role struct {
	ID          uint64        `json:"" sql:"index"`
	Name        string        `json:"name,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty" gorm:"many2many:roles_permissions;"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// Permission struct
type Permission struct {
	ID        uint64    `json:"-" sql:"index"`
	Name      string    `json:"name,omitempty"`
	Role      []*Role   `json:"-" gorm:"many2many:roles_permissions;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
