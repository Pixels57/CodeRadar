package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []User

	if err := h.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
