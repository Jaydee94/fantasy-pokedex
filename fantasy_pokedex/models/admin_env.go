package models

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Call this after migration and before DB is used
func EnsureAdminFromEnv(db *gorm.DB) {
	pw := os.Getenv("ADMIN_PASSWORD")
	if pw == "" {
		return
	}
	var count int64
	db.Model(&Admin{}).Count(&count)
	if count > 0 {
		return // Admin already exists
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[admin] Failed to hash ADMIN_PASSWORD: %v", err)
		return
	}
	admin := Admin{Password: string(hash)}
	if err := db.Create(&admin).Error; err != nil {
		log.Printf("[admin] Failed to create admin from env: %v", err)
	} else {
		log.Printf("[admin] Admin password from environment set.")
	}
}
