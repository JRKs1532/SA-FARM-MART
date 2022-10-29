package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/JRKs1532/sa-65-example/entity"
)

// POST /positions
func CreateEducation(c *gin.Context) {
	var education entity.Education
	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&education).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": education})
}

// GET /position/:id
func GetEducation(c *gin.Context) {
	var education entity.Education
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&education); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "education not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": education})
}

// GET /positions
func ListEducations(c *gin.Context) {
	var  educations []entity.Education
	if err := entity.DB().Raw("SELECT * FROM educations").Scan(&educations).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": educations})
}

// DELETE /positions/:id
func DeleteEducation(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM educations WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "education not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /positions
func UpdateEducation(c *gin.Context) {
	var education entity.Education
	if err := c.ShouldBindJSON(&education); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", education.ID).First(&education); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "education not found"})
		return
	}

	if err := entity.DB().Save(&education).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": education})
}