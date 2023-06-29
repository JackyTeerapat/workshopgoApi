package student

import (
	"net/http"

	"workshopgoApi/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StudentHandler struct {
	db *gorm.DB
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{db: db}
}

func (h *StudentHandler) GetAllStudent(c echo.Context) error {
	students := []models.Student{}

	result := h.db.Find(&students)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "student not found")
		return result.Error
	}

	return c.JSON(http.StatusOK, students)
}
func (h *StudentHandler) GetStudentById(c echo.Context) error {
	id := c.Param("id")

	student := models.Student{}

	if result := h.db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, "student not found")
	}

	return c.JSON(http.StatusOK, student)
}
func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var student models.Student

	if err := c.Bind(&student); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data format")
	}

	if result := h.db.Create(&student); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create student")
	}

	return c.JSON(http.StatusCreated, &student)
}
func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	id := c.Param("id")

	var student models.Student
	if result := h.db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, "Student not found")
	}

	var updatedStudent models.Student
	if err := c.Bind(&updatedStudent); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data format")
	}

	if result := h.db.Model(&student).Updates(updatedStudent); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update student")
	}

	return c.JSON(http.StatusOK, &student)
}
func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	var student models.Student
	if result := h.db.First(&student, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, "Student not found")
	}

	if result := h.db.Delete(&student); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON(http.StatusOK, "Student deleted successfully")
}
