package services

import (
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/forms"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"

	"errors"

)

type CourseService struct{}

func (svc CourseService) CreateCourse(user models.User, courseForm forms.CreateCourseForm) (course models.Course, err error) {
	//token에서 user의 정보를 가져오고, 교수인지 확인한다(Isprofessor ==1)
	professorid := user.ID
	if !user.Isprofessor {
		return course, errors.New("당신은 교수가 아닙니다")
	}

	//가져온 데이터를 이용해 새 course를 생성한다
	course = models.Course{Coursenumber: courseForm.Coursenumber, Coursename: courseForm.Coursename}
	result := initializers.DB.Create(&course)	
	if result.Error != nil {
		return course, errors.New("강의를 생성할 수 없습니다")
	}

	//가져온 데이터를 이용해 새 teach를 생성한다
	teach := models.Teach{ProfessorID: professorid, CourseID: course.ID}
	result = initializers.DB.Create(&teach)
	if result.Error != nil {
		return course, errors.New("Failed to create teach")
	}

	return course, nil

}

func (svc CourseService) SearchAllCourse() (courses []models.Course, err error) {
	//모든 course를 가져온다
	result := initializers.DB.Find(&courses)
	if result.Error != nil {
		return courses, errors.New("강의를 찾을 수 없습니다")
	}

	if len(courses) == 0 {
		return courses, errors.New("등록된 강의가 없습니다")
	}

	return courses, nil
}

func (svc CourseService) SearchTeachCourse(user models.User) (courses []models.Course, err error) {
	//token에서 user의 정보를 가져온다
	professorid := user.ID

	//user가 professor인지 확인한다(Isprofessor ==1)
	if !user.Isprofessor {
		return courses, errors.New("당신은 교수가 아닙니다")
	}

	//가져온 데이터를 이용해 teach에서 course를 찾는다
	var teaches []models.Teach
	result := initializers.DB.Where("professor_id = ?", professorid).Find(&teaches)
	if result.Error != nil {
		return courses, errors.New("Failed to find teaches")
	}

	//찾은 course들을 반환한다
	for _, teach := range teaches {
		var course models.Course
		result := initializers.DB.Where("id = ?", teach.CourseID).Find(&course)
		if result.Error != nil {
			return courses, errors.New("Failed to find course")
		}
		courses = append(courses, course)
	}

	if len(courses) == 0 {
		return courses, errors.New("강의를 찾을 수 없습니다")
	}

	return courses, nil
}

func (svc CourseService) SearchApplyCourse(user models.User) (courses []models.Course, err error) {
	//token에서 user의 정보를 가져온다
	studentid := user.ID

	//user가 student인지 확인한다(Isprofessor ==0)
	if user.Isprofessor {
		return courses, errors.New("당신은 학생이 아닙니다")
	}

	//가져온 데이터를 이용해 apply에서 course를 찾는다
	var apply []models.Apply
	result := initializers.DB.Where("student_id = ?", studentid).Find(&apply)
	if result.Error != nil {
		return courses, errors.New("Failed to find apply")
	}

	//찾은 course를 반환한다
	for _, apply := range apply {
		var course models.Course
		result := initializers.DB.Where("id = ?", apply.CourseID).Find(&course)
		if result.Error != nil {
			return courses, errors.New("Failed to find course")
		}
		courses = append(courses, course)
	}


	if len(courses) == 0 {
		return courses, errors.New("수강한 강의가 없습니다")
	}

	return courses, nil
}