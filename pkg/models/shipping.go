package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShippingCost struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type ShippingAddress struct {
	ID          string `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID      string `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
}

func (s *ShippingAddress) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}

type ShippingAddressRequest struct {
	Name        string `json:"name" validate:"required string"`
	PhoneNumber string `json:"phone_number" validate:"required string"`
	Address     string `json:"address" validate:"required string"`
	City        string `json:"city" validate:"required string"`
}

type ShippingAddressResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
}
