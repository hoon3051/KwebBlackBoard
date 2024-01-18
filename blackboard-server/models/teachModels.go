package models

import "gorm.io/gorm"

//교수가 강의하는 강의들
type Teach struct {
	gorm.Model
	CourseID    uint
	ProfessorID uint

	//teach는 course의 id를 CourseID로, user의 id를 professorID로 ManyToOne 매핑된다
	User   User   `gorm:"foreignkey:ProfessorID"`
	Course Course `gorm:"foreignkey:CourseID"`
	
}
