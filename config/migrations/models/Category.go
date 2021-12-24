package models

import "time"

type Category struct {
	ID            uint    `gorm:"primaryKey"`
	Name          string  `gorm:"type:varchar(20);not null"`
	Description   string  `gorm:"type:varchar(200);not null"`
	Interest_rate float64 `gorm:"not null"`
	Products      []Product
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
