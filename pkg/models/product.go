package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	ID   string `json:"id" gorm:"primary_key , type:uid, default:uuid_generate_v4()"`
	Name string `json:"product_name" validate:"required string"`
	// Size is the size of the product contain S, M, L
	Sizes      datatypes.JSON `json:"sizes" validate:"required"`
	Details    string         `json:"product_details" validate:"required string"`
	Price      float64        `json:"price" validate:"required float"`
	ImageURL   string         `json:"image_url" validate:"required string"`
	Condition  string         `json:"condition" validate:"required string"`
	CategoryID string         `json:"category_id" validate:"required string"`
	Category   Category       `json:"category" gorm:"foreignKey:CategoryID"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}

type ProductRequest struct {
	Name       string  `json:"product_name" validate:"required string"`
	Details    string  `json:"product_details" validate:"required string"`
	Price      float64 `json:"price" validate:"required float"`
	ImageURL   string  `json:"image_url" validate:"required string"`
	Condition  string  `json:"condition" validate:"required string"`
	CategoryID string  `json:"category_id" validate:"required string"`
}
