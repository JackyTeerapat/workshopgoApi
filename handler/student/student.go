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

	//find เป็น คำส่ง SELECT * FROM ตัวที่มีค่าตรงกับ model ที่ใส่ไป
	// err:= h.db.Find(&ตัวแปรที่ต้องการใส่ค่า) โดยจะ return err เป็น err.Error
	if ??? :=??? != nil {
		c.JSON(http.StatusBadRequest, "student not found")
		return ???.Error
	}
	
	//c.JSON(statusCode int, i interface{})
	// c.json จะเป็น JSON response โดย parameter แรกจะเป็น int (status code) อันที่ 2 คือใส่อะไรไปจะ return อันนั้น
	//200: OK - ร้องขอสำเร็จและมีข้อมูลที่ส่งกลับ
	// 201: Created - การสร้างข้อมูลสำเร็จ
	// 400: Bad Request - การร้องขอไม่ถูกต้องหรือมีข้อมูลไม่ครบถ้วน
	// 401: Unauthorized - การร้องขอไม่ได้รับการรับรองตัวตน
	// 403: Forbidden - ไม่ได้รับอนุญาตให้เข้าถึงทรัพยากร
	// 404: Not Found - ทรัพยากรที่ร้องขอไม่พบ
	// 500: Internal Server Error - เกิดข้อผิดพลาดภายในเซิร์ฟเวอร์ 

	return c.JSON(???,???)
}
func (h *StudentHandler) GetStudentById(c echo.Context) error {
	id := c.Param("id")

	student := models.Student{}

	//First เป็น คำส่ง SELECT * FROM ตัวที่มีค่าตรงกับ model ที่ใส่ไป คล้าย find แต่ parameter ตัวที่สองคือ where
	// err:=h.db.First(&ตัวแปรที่ต้องการใส่ค่า, ตัวที่ต้องการ หา) โดยจะ return err เป็น err.Error
	if ??? := ???; ??? != nil {
		return c.JSON(http.StatusNotFound, "student not found")
	}

	return c.JSON(???, ???)
}
func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var student models.Student

	// c.Bind(&ตัวแปรที่ต้องการ map ค่า) เป็น function เพื่อรับค่าจาก request มา map
	if err := c.Bind(&????); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid data format")
	}
	//err:= h.db.Create (& ตัวที่ต้องการ insert)คือการ insert ข้อมูลตาม struct โดยจะ return err เป็น err.Error
	// INSERT INTO student ("name", "gender", "course") VALUES (?, ?, ?)

	if ??? := h.db.Create(&???); ???.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create student")
	}

	return c.JSON(???, &???)
}
func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	id := c.Param("id")


	//หานักเรียนโดย id แล้ว map เก็บไว้ พร้อมดัก errors hint.GetStudentById
	

	//รับค่า request มาเก็บเป็นตัวแปรพร้อมดัก errors hint.CreateStudent
	

	//Syntax ในการ update
	//err:= h.db.Model(&model).Updates(ค่าที่ request ที่ map มา) โดยจะ return err เป็น err.Error
	if ??? := ???; ???.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update student")
	}

	return c.JSON(???, ???)
}
func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	//หานักเรียนโดย id แล้ว map เก็บไว้ พร้อมดัก errors hint.GetStudentById

	//Syntax ในการลบ
	//err:= h.db.Delete(&model ทีต้องการลบ) โดยจะ return err เป็น err.Error
	if (???, := (???,; (???,.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON((???,, ???)
}

type StudentWithTeacherName struct {
	models.Student
	TeacherName string `json:"teacher_name"`
}

func (h *StudentHandler) GetAllStudentsWithTeacherName(c echo.Context) error {
	var students []StudentWithTeacherName

	
	query := `
		SELECT student.*, teacher.name AS teacher_name
		FROM student
		INNER JOIN teacher ON student.course = teacher.coursename
	`

		//err:=h.db.Raw(query).Scan(&ตัวแปรที่ต้องการใส่ค่า) พร้อมดัก error โดยจะ return err เป็น err.Error
		if ??? := ???

	return c.JSON(???, ???)
}
