package models

import "gorm.io/gorm"

//사용자 정보들
type User struct {
	gorm.Model
	Username      string
	Password      string
	Displayname   string
	Studentnumber int32
	Isprofessor   bool

	//user의 id는 teach, apply에 OneToMany 매핑된다
	Teach []Teach `gorm:"foreignkey:ProfessorID"`
	Apply []Apply `gorm:"foreignkey:StudentID"`
}
