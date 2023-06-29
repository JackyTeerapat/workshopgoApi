package models

type Student struct {
	ID     int    `gorm:"primary_key;default:nextval('student_studentid_seq'::regclass);not null"`
	Name   string `gorm:"type:varchar(255);column:name"`
	Gender string `gorm:"type:varchar(255);not null"`
	Course string `gorm:"type:varchar(255);not null"`
}

func (Student) TableName() string {
	return "student"
}
