package controllers

import (
	"net/http"

	"github.com/hoon3051/KwebBlackBoard/blackboard-server/forms"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/services"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/utils"

	"github.com/gin-gonic/gin"
)

func CreateBoard(c *gin.Context) {
	//req body로부터 course의 정보를 가져온다
	var CreateBoardForm forms.CreateBoardForm
	if err := c.ShouldBindJSON(&CreateBoardForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validation을 거친다
	boardForm := forms.BoardForm{}
	if err := boardForm.Create(CreateBoardForm); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//parameter가 string이므로 uint로 변환하는 과정을 거친다
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	//교수가 가르치는 과목인지 확인한다
	teachService := services.TeachService{}
	_, err = teachService.GetTeach(*user, courseid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//가져온 데이터를 이용해 새 board를 생성한다
	boardService := services.BoardService{}
	board, err := boardService.CreateBoard(CreateBoardForm, courseid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created board",
		"board":   board,
	})

}

func SearchAllBoard(c *gin.Context) {
	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//user가 apply한 course들을 가져온다
	courseService := services.CourseService{}
	courses, err := courseService.SearchApplyCourse(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//course들의 board들을 가져온다
	boardService := services.BoardService{}
	boards, err := boardService.GetAllBoards([]models.Course{courses})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"boards": boards,
	})
}

func SearchCourseBoard(c *gin.Context) {
	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//parameter가 string이므로 uint로 변환하는 과정을 거친다
	courseID, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	//user가 course에 대한 권한이 있는지 확인한다
	if user.Isprofessor {
		//user가 교수라면 teach를 이용해서 course를 가르치는지 확인한다
		teachService := services.TeachService{}
		_, err := teachService.GetTeach(*user, courseID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		//user가 학생이라면 apply를 이용해서 course를 등록했는지 확인한다
		applyService :=services.ApplyService{}
		_, err := applyService.GetApply(*user, courseID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	//courseID로 board들을 가져온다
	boardService := services.BoardService{}
	boards, err := boardService.GetBoards(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"boards": boards,
	})

}

func GetBoard(c *gin.Context) {
	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//parameter가 string이므로 uint로 변환하는 과정을 거친다
	boardid, err := utils.GetUintParam(c, "board_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid board ID",
		})
		return
	}

	//boardID로 courseID를 가져온다
	boardService := services.BoardService{}
	courseID, err := boardService.GetCourseIDofBoard(boardid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//user가 board에 대한 권한이 있는지 확인한다
	if user.Isprofessor {
		//user가 교수라면 teach를 이용해서 course를 가르치는지 확인한다
		teachService := services.TeachService{}
		_, err := teachService.GetTeach(*user, courseID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		//user가 학생이라면 apply를 이용해서 course를 등록했는지 확인한다
		applyService :=services.ApplyService{}
		_, err := applyService.GetApply(*user, courseID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	//boardID로 board를 가져온다
	board, err := boardService.GetBoard(*user, boardid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//응답해준다
	c.JSON(http.StatusOK, gin.H{
		"board": board,
	})

}
