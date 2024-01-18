package models

import "gorm.io/gorm"

//수업들
type Course struct {
	gorm.Model
	Coursenumber string
	Coursename   string

	//course의 id는 teach, apply, board에 OneToMany 매핑된다
	Teach []Teach `gorm:"foreignkey:CourseID"`
	Apply []Apply `gorm:"foreignkey:CourseID"`
	Board []Board `gorm:"foreignkey:CourseID"`
}
