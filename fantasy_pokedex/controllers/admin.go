package controllers

import (
	"fantasy_pokedex/config"
	"fantasy_pokedex/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminPasswordRequest struct {
	Password string `json:"password"`
}

type AdminPasswordSetResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AdminPasswordStatusResponse struct {
	Set bool `json:"set"`
}

// GET /admin/password-set
func AdminPasswordSetStatus(c *gin.Context) {
	var admin models.Admin
	err := config.DB.First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, AdminPasswordStatusResponse{Set: false})
		return
	}
	c.JSON(http.StatusOK, AdminPasswordStatusResponse{Set: true})
}

// POST /admin/set-password
func AdminSetPassword(c *gin.Context) {
	var admin models.Admin
	err := config.DB.First(&admin).Error
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusForbidden, gin.H{"error": "Password already set"})
		return
	}
	var req AdminPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	admin.Password = string(hash)
	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving password"})
		return
	}
	c.JSON(http.StatusOK, AdminPasswordSetResponse{Success: true, Message: "Password set successfully"})
}

// POST /admin/login
func AdminLogin(c *gin.Context) {
	var req AdminPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var admin models.Admin
	err := config.DB.First(&admin).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusForbidden, gin.H{"error": "No admin password set"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	c.JSON(http.StatusOK, AdminPasswordSetResponse{Success: true, Message: "Login successful"})
}
