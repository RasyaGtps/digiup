package dao

import (
	"time"

	"gorm.io/gorm"
)

type Borrowing struct {
	ID         uint       `gorm:"primaryKey"`
	BorrowDate time.Time  `gorm:"not null;"`
	ReturnDate *time.Time `gorm:""`
	BookID     uint       `gorm:"not null;"`
	PersonID   uint       `gorm:"not null"`
	Book       Book       `gorm:"foreignKey:BookID"`
	Person     Person     `gorm:"foreignKey:PersonID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
