package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCart struct {
	ID        string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ProductID string  `json:"item_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Sizes     string  `json:"sizes"`
}

func (u *UserCart) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserCartRequest struct {
	ProductID string `json:"item_id" validate:"required string"`
	Quantity  int    `json:"quantity" validate:"required int"`
	Sizes     string `json:"sizes" validate:"required string"`
}
