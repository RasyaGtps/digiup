package dao

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        uint      `gorm:"primarykey;type:bigint"`
	FullName  string    `gorm:"type:varchar(56);not null"`
	Gender    string    `gorm:"type:enum('m','f');default:null"`
	BirthDate time.Time `gorm:"type:datetime;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
