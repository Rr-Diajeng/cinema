package model

type Class struct {
	ID            uint    `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Name          string  `gorm:"type:varchar;not_null"`
	PriceModifier float32 `gorm:"type:float"`
	Seats         []Seats `gorm:"foreignKey:ClassID"`
}