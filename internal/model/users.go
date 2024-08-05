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