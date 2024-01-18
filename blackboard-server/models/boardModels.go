package models

import "gorm.io/gorm"

//수업의 게시판들
type Board struct {
	gorm.Model
	CourseID uint
	Title    string
	Desc     string

	//board는 course의 id를 CourseID로 ManyToOne 매핑된다
	Course Course `gorm:"foreignkey:CourseID"`
}
