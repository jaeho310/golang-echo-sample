package model

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name,omitempty"`
	Cards     []Card     `gorm:"foreignKey:UserId"json:"cards"`
}
