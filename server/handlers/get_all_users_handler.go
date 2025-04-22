package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Image    string   `json:"image"`
	Skills   []string `json:"skills"`
	JobTitle string   `json:"jobtitle"`
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert users to response format with base64 images
	responseUsers := make([]UserResponse, len(users))
	for i, user := range users {
		// Parse skills from JSON string
		var skills []string
		if err := json.Unmarshal([]byte(user.Skills), &skills); err != nil {
			skills = []string{}
		}

		// Convert image to base64 if it exists
		var base64Image string
		if user.ImagePath != "" {
			imageData, err := ioutil.ReadFile(user.ImagePath)
			if err == nil {
				base64Image = base64.StdEncoding.EncodeToString(imageData)
			}
		}

		responseUsers[i] = UserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Image:    base64Image,
			Skills:   skills,
			JobTitle: user.JobTitle,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"users": responseUsers,
	})
}
