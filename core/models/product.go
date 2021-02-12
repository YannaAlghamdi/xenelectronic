package models

import (
	uuid "github.com/satori/go.uuid"
)

// type BaseModel struct {
// 	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
// 	CreatedAt time.Time  `json:"created_at"`
// 	UpdatedAt time.Time  `json:"updated_at"`
// 	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
// }

type Product struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
