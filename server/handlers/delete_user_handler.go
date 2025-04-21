package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	var user User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found	"})
		return
	}

	if user.ImagePath != "" {
		if err := os.Remove(user.ImagePath); err != nil {
			log.Printf("Failed to delete image: %v", err)
		}
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "User deleted successfully"})
}
