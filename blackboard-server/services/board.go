package services

import (
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/forms"
)

type BoardService struct{}

func (svc BoardService) CreateBoard(boardForm forms.CreateBoardForm, courseID uint) (board models.Board, err error) {
	//가져온 데이터를 이용해 새 board를 생성한다
	board = models.Board{CourseID: courseID, Title: boardForm.Title, Content: boardForm.Content}
	result := initializers.DB.Create(&board)
	if result.Error != nil {
		return board, result.Error
	}

	return board, nil
}

func (svc BoardService) GetBoard(user models.User, boardID uint) (board models.Board, err error) {
	//user가 교수면 teach를 이용하여 course를 가져오고, course의 board중 boardID와 일치하는 board를 가져온다


	//boardID로 board를 가져온다
	board = models.Board{}
	result := initializers.DB.First(&board, boardID)
	if result.Error != nil {
		return board, result.Error
	}

	return board, nil
}

func (svc BoardService) GetBoards(courseID uint) (boards []models.Board, err error) {
	//courseID로 board를 가져온다
	result := initializers.DB.Where("course_id = ?", courseID).Find(&boards)
	if result.Error != nil {
		return boards, result.Error
	}

	return boards, nil
}

func (svc BoardService) GetCourseIDofBoard(boardID uint) (courseID uint, err error) {
	//boardID로 board를 가져온다
	board := models.Board{}
	result := initializers.DB.First(&board, boardID)
	if result.Error != nil {
		return courseID, result.Error
	}

	return board.CourseID, nil
}

func (svc BoardService) GetAllBoards(courses []models.Course) (boards []models.Board, err error) {
	//courses를 이용해 board들을 가져온다
	var courseIDs []uint
	for _, course := range courses {
		courseIDs = append(courseIDs, course.ID)
	}

	result := initializers.DB.Where("course_id IN ?", courseIDs).
	Where("updated_at DESC").
	Limit(10).
	Find(&boards)
	if result.Error != nil {
		return boards, result.Error
	}

	return boards, nil
}
