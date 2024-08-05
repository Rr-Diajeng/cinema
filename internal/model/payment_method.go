package model

type PaymentMethods struct {
	ID           uint           `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	Name         string         `gorm:"type:varchar;not_null"`
	Transactions []Transactions `gorm:"foreignKey:PaymentMethodsID"`
}