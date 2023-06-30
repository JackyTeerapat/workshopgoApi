package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"workshopgoApi/handler/student"
)

var dsn = "postgres://sfoernga:c5qveU2H3emvMxwFhlKWcUC12pQNN438@satao.db.elephantsql.com/sfoernga"
var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

func main() {
	if err != nil {
		panic("error to connect to database")
	}
	e := echo.New()

	studentHandler := student.NewStudentHandler(db)

	e.GET("/student", studentHandler.GetAllStudent)
	e.POST("/student", studentHandler.CreateStudent)

	e.GET("/student/:id", studentHandler.GetStudentById)
	e.PATCH("/student/:id", studentHandler.UpdateStudent)
	e.DELETE("/student/:id", studentHandler.DeleteStudent)

	e.GET("/studentWithTeacher", studentHandler.GetAllStudentsWithTeacherName)
	e.Logger.Fatal(e.Start(":9000"))
}

//เป็นพื้นฐานการทำ CRUD ของ GO
// ทำการเติม Function ต่างใน Folder Handler/student ให้สามารถใช้งานได้ โดยแทนค่าข้าไปใน เครื่องหมาย ??? ที่มีไว้ให้
