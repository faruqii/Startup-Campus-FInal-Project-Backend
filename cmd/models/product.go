package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        string  `json:"id" gorm:"primary_key , type:uid, default:uuid_generate_v4()"`
	Category  string  `json:"category" validate:"required string"`
	Price     float64 `json:"price" validate:"required float"`
	Condition string  `json:"condition" validate:"required string"`
	Name      string  `json:"product_name" validate:"required string"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}

type ProductRequest struct {
	Category  string  `json:"category" validate:"required string"`
	Price     float64 `json:"price" validate:"required float"`
	Condition string  `json:"condition" validate:"required string"`
	Name      string  `json:"product_name" validate:"required string"`
}
