package model

import "gorm.io/gorm"

type TransactionStatus string

const (
	Pending   TransactionStatus = "pending"
	Ongoing   TransactionStatus = "ongoing"
	Completed TransactionStatus = "completed"
	Cancelled TransactionStatus = "cancelled"
)

type Transactions struct {
	gorm.Model
	ID                uint    `gorm:"type:bigint;primary_key,AUTO_INCREMENT"`
	TotalPrice        float32 `gorm:"type:varchar"`
	PaymentMethodsID  uint
	TransactionStatus TransactionStatus `gorm:"type:transaction_status;not_null"`
	Tickets           []Tickets         `gorm:"foreignKey:TransactionID"`
	PaymentMethods    PaymentMethods    `gorm:"foreignKey:PaymentMethodsID"`
}