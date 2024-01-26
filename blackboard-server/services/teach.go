package services

import (
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"

	"errors"
)

type TeachService struct{}

func (svc TeachService) GetTeach(user models.User, courseid uint) (teach models.Teach, err error) {
	//token에서 user의 정보를 가져온다
	professorid := user.ID

	//user가 professor인지 확인한다(Isprofessor ==1)
	if !user.Isprofessor {
		return teach, errors.New("You are not a professor")
	}

	//가져온 데이터를 이용해 teach를 찾는다
	teach = models.Teach{ProfessorID: professorid, CourseID: courseid}
	result := initializers.DB.Where(&teach).First(&teach)
	if result.Error != nil {
		return teach, result.Error
	}

	return teach, nil

}