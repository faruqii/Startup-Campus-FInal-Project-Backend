package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   string `json:"id" gorm:"primary_key , type:uid, default:uuid_generate_v4()"`
	Name string `json:"name" validate:"required string"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return
}

type CategoryRequest struct {
	Name string `json:"name" validate:"required string"`
}

// Path: pkg\models\category.go
