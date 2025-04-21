package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GetUsersBySkill(c *gin.Context) {
	skill := strings.ToLower(c.Param("skill"))
	var users []User

	if err := h.DB.Where("item_list LIKE ?", "%"+skill+"%").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
