package model

import "time"

type Card struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	UserId    uint       `json:"userId"`
	User      User
}
