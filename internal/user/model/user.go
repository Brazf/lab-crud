package model

import "time"

const (
	RoleRead  = "READ"
	RoleWrite = "WRITE"
	RoleRoot  = "ROOT"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type Organization struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type OrganizationUser struct {
	OrganizationID uint   `gorm:"primaryKey" json:"organization_id"`
	UserID         uint   `gorm:"primaryKey" json:"user_id"`
	Role           string `json:"role"`
}
