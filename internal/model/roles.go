package model

type Roles struct {
	ID    uint    `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Name  string  `gorm:"type:varchar;size:100;not_null"`
	Users []Users `gorm:"foreignKey:RoleID"`
}