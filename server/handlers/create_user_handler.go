package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"not null" json:"email"`
	ImagePath string `gorm:"column:imagepath" json:"imagepath"`
	Skills    string `gorm:"column:item_list;type:json" json:"skills"`
	JobTitle  string `gorm:"column:jobtitle" json:"jobtitle"`
}

type CreateUserRequest struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	JobTitle string   `json:"jobTitle"`
	Image    string   `json:"image"`
	Skills   []string `json:"skills"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert skills array to JSON string
	var skillsJSON []byte
	var err error
	if len(req.Skills) == 0 {
		skillsJSON = []byte("[]") // Empty array as JSON
	} else {
		skillsJSON, err = json.Marshal(req.Skills)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Create user in database
	newUser := User{
		Name:     req.Name,
		Email:    req.Email,
		JobTitle: req.JobTitle,
		Skills:   string(skillsJSON),
	}

	result := h.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Handle base64 image if provided
	if req.Image != "" {
		// Remove the "data:image/png;base64," prefix if present
		base64Data := req.Image
		if strings.Contains(req.Image, "base64,") {
			base64Data = strings.Split(req.Image, "base64,")[1]
		}

		// Decode base64 string
		imageData, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 image data"})
			return
		}

		// Save image to file
		filename := fmt.Sprintf("user_%d.png", newUser.ID)
		filepath := filepath.Join(uploadsDir, filename)

		if err := ioutil.WriteFile(filepath, imageData, 0644); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update user with image path
		newUser.ImagePath = filepath
		h.DB.Save(&newUser)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
}
