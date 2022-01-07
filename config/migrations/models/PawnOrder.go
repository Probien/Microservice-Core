package models

import "time"

type PawnOrder struct {
	ID            uint `gorm:"primaryKey"`
	EmployeeID    uint `gorm:"not null"`
	CustomerID    uint `gorm:"not null"`
	StatusID      uint
	TotalMount    float64 `gorm:"not null"`
	Monthly       bool    `gorm:"not null"`
	Products      []Product
	Endorsements  []Endorsement
	CutOffDay     time.Time
	ExtensionDate time.Time
}