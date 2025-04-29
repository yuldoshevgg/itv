package movie

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Title     string         `json:"title" binding:"required"`
	Director  string         `json:"director" binding:"required"`
	Year      int            `json:"year" binding:"required"`
	Plot      string         `json:"plot"`
}
