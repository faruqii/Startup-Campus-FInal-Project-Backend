package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID                string          `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID            string          `json:"user_id"`
	User              User            `json:"user" gorm:"foreignKey:UserID"`
	ShippingMethod    string          `json:"shipping_method"`
	ShippingAddressID string          `json:"shipping_address_id"`
	ShippingAddress   ShippingAddress `json:"shipping_address" gorm:"foreignKey:ShippingAddressID"`
	ShippingPrice     int64           `json:"shipping_price"`
	TotalPrice        float64         `json:"total_price"`
	CreatedAt         time.Time       `json:"created_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return
}

type OrderRequest struct {
	ShippingMethod  string `json:"shipping_method" validate:"required string"`
	ShippingAddress string `json:"shipping_address" validate:"required string"`
	ShippingPrice   int64  `json:"shipping_price" validate:"required float"`
}

type OrderResponse struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UserID    string    `json:"user_id"`
	UserEmail string    `json:"user_email"`
	Total     float64   `json:"total"`
}
