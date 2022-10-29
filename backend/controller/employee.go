package controller

import (
	"net/http"

	"github.com/JRKs1532/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func CreateEmployees(c *gin.Context) {

	var employee entity.Employee
	var education entity.Education
	var gedner entity.Gender
	var position entity.Position

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา gedner ด้วย id
	if tx := entity.DB().Where("id = ?", employee.GenderID).First(&gedner); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gedner not found"})
		return
	}

	// 10: ค้นหา education ด้วย id
	if tx := entity.DB().Where("id = ?", employee.EducationID).First(&education); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "education not found"})
		return
	}

	// 10: ค้นหา position ด้วย id
	if tx := entity.DB().Where("id = ?",employee.PositionID).First(&position); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "position not found"})
		return
	}

	// 12: สร้าง Employee
	wv := entity.Employee{
		FirstName:    	employee.FirstName,     // ตั้งค่าฟิลด์
		LastName:      	employee.LastName,      // ตั้งค่าฟิลด์
		Telephone:     	employee.Telephone,     // ตั้งค่าฟิลด์
		Email:			employee.Email,
		Slary:			employee.Slary,
		Password:		employee.Password,
		Education:  	education,             // โยงความสัมพันธ์กับ Entity education
		Position:  		position,             // โยงความสัมพันธ์กับ Entity position
		Gender:       	gedner,                  // โยงความสัมพันธ์กับ Entity gender
		StartJob: 		employee.StartJob, // ตั้งค่าฟิลด์ 

	// 13: บันทึก
	}
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /employee/:id
func GetEmployee(c *gin.Context) {
	var employee entity.Employee
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// GET /employees
func ListEmployees(c *gin.Context) {
	var employees []entity.Employee
	if err := entity.DB().Preload("Education").Preload("Position").Preload("Gender").Raw("SELECT * FROM employees").Find(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// DELETE /employees/:id
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM employees WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /employees
func UpdateEmployeeIn(c *gin.Context) {
	var employee entity.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", employee.ID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	if err := entity.DB().Save(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}