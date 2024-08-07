package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID       uint   `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Username string `gorm:"type:varchar;size:100;not_null"`
	Email    string `gorm:"type:varchar;not_null;unique"`
	Password string `gorm:"type:varchar;not_null"`
	RoleID   uint
	Role     Roles     `gorm:"foreignKey:RoleID"`
	Tickets  []Tickets `gorm:"foreignKey:UserID"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"roleId"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type UserProfileResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"roleName"`
}
