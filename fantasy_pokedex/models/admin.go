package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	ID       uint   `gorm:"primaryKey"`
	Password string `gorm:"not null"`
}

func MigrateAdmin(db *gorm.DB) error {
	return db.AutoMigrate(&Admin{})
}
