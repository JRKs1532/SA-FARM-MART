package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JRKs1532/sa-65-example/entity"
)

// POST /positions
func CreatePosition(c *gin.Context) {
	var position entity.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&position).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": position})
}

// GET /position/:id
func GetPosition(c *gin.Context) {
	var position entity.Position
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&position); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": position})
}

// GET /positions
func ListPositions(c *gin.Context) {
	var positions []entity.Position
	if err := entity.DB().Raw("SELECT * FROM positions").Scan(&positions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": positions})
}

// DELETE /positions/:id
func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM positions WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /positions
func UpdatePosition(c *gin.Context) {
	var position entity.Position
	if err := c.ShouldBindJSON(&position); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", position.ID).First(&position); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rposition not found"})
		return
	}

	if err := entity.DB().Save(&position).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": position})
}