package models

import "github.com/jinzhu/gorm"

type EventOrder struct {
	gorm.Model
	OrderToken    string `json:"order_token"`
	PaymentStatus string `json:"payment_status"`
	PaymentMethod string `json:"payment_method"`
	MasterEventID uint
}
