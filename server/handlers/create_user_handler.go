package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Email     string `gorm:"not null" json:"email"`
	ImagePath string `gorm:"column:imagepath" json:"imagepath"`
	ItemList  string `gorm:"type:json" json:"item_list"`
	JobTitle  string `gorm:"column:jobtitle" json:"jobtitle"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {

	uploadsDir := "./uploads"
	if err := os.MkdirAll(uploadsDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get other form data
	name := c.PostForm("name")
	email := c.PostForm("email")
	jobtitle := c.PostForm("jobtitle")
	image := c.PostForm("imagepath")
	itemList := c.PostForm("item_list")

	// Create user in database
	newUser := User{
		Name:     name,
		Email:    email,
		JobTitle: jobtitle,
		ItemList: itemList,
	}

	result := h.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if image != "" {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
			return
		}

		filename := fmt.Sprintf("user_%d%s", newUser.ID, filepath.Ext(file.Filename))
		filepath := filepath.Join(uploadsDir, filename)

		if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newUser.ImagePath = filepath
		h.DB.Save(&newUser)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": newUser})
}
