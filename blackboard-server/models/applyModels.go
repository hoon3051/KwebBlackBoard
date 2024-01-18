package models

import "gorm.io/gorm"

//학생들이 수강한 강의들
type Apply struct {
	gorm.Model
	CourseID  uint
	StudentID uint

	//apply는 course의 id를 CourseID로, user의 id를 studentID로 ManyToOne 매핑된다
	User   User   `gorm:"foreignkey:StudentID"`
	Course Course `gorm:"foreignkey:CourseID"`
}
