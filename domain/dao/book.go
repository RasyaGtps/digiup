package dao

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:56;"`
	Subtitle      *string   `gorm:"size:64;"`
	PublisherID   uint      `gorm:"not null;"`
	AuthorID      uint      `gorm:"not null"`
	BookPublisher Publisher `gorm:"foreignKey:PublisherID;"`
	BookAuthor    Author    `gorm:"foreignKey:AuthorID;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
