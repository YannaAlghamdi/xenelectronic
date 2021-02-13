package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        string     `json:"id,omitempty" sql:"type:uuid;primary_key"`
	CreatedAt time.Time  `json:"datetime_created,omitempty" gorm:"type:timestamp"`
	UpdatedAt time.Time  `json:"datetime_last_updated,omitempty" gorm:"type:timestamp"`
	DeletedAt *time.Time `json:"datetime_deleted,omitempty" gorm:"type:timestamp"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New().String())
}
