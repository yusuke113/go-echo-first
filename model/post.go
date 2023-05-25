package model

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Body      string    `json:"body" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user" gorm:"foreignKey:UserID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Body      string    `json:"body" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
