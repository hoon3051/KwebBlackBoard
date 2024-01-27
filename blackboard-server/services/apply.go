package services

import (
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"

	"errors"
)

type ApplyService struct{}

func (svc ApplyService) CreateApply(user models.User, courseID uint) (apply models.Apply, err error) {
	//token에서 user의 정보를 가져온다
	studentid := user.ID

	//user가 student인지 확인한다(Isprofessor ==0)
	if user.Isprofessor {
		return apply, errors.New("You are not a student")
	}

	//가져온 데이터를 이용해 이미 등록한 apply가 있는지 확인한다
	apply = models.Apply{StudentID: studentid, CourseID: courseID}
	result := initializers.DB.Where(&apply).First(&apply)
	if result.Error == nil {
		return apply, errors.New("You already applied")
	}

	//가져온 데이터를 이용해 새 apply를 생성한다
	result = initializers.DB.Create(&apply)
	if result.Error != nil {
		return apply, result.Error
	}

	return apply, nil

}

func (svc ApplyService) GetAppliedStudent(courseID uint) (students []models.User, err error) {
	//가져온 courseID와 일치하는 CourseID를 가진 apply들을 찾는다
	var applies []models.Apply
	result := initializers.DB.Where("course_id = ?", courseID).Find(&applies)
	if result.Error != nil {
		return students, result.Error
	}

	//찾은 apply들의 studentID를 이용해 user들을 찾는다
	var studentsID []uint
	for _, apply := range applies {
		studentsID = append(studentsID, apply.StudentID)
	}

	result = initializers.DB.Where("id IN ?", studentsID).Find(&students)
	if result.Error != nil {
		return students, result.Error
	}

	if len(students) == 0 {
		return students, errors.New("There is no student")
	}

	return students, nil
	
}

func (svc ApplyService) DeleteApply(user models.User, studentID uint, courseID uint) (err error) {
	//user가 student인지 확인한다(Isprofessor ==0)
	if !user.Isprofessor {
		return errors.New("You are not a professor")
	}

	//가져온 데이터를 이용해 apply를 찾는다
	apply := models.Apply{StudentID: studentID, CourseID: courseID}
	result := initializers.DB.Where(&apply).First(&apply)
	if result.Error != nil {
		return result.Error
	}

	//찾은 apply를 삭제한다
	result = initializers.DB.Delete(&apply)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (svc ApplyService) GetApply(user models.User, courseID uint) (apply models.Apply, err error) {
	//user가 student인지 확인한다(Isprofessor ==0)
	if user.Isprofessor {
		return apply, errors.New("You are not a student")
	}

	//가져온 데이터를 이용해 apply를 찾는다
	apply = models.Apply{StudentID: user.ID, CourseID: courseID}
	result := initializers.DB.Where(&apply).First(&apply)
	if result.Error != nil {
		return apply, result.Error
	}

	return apply, nil

}