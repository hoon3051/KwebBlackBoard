package board

import (
	"hoon/KwebBlackBoard/blackboard-server/initializers"
	"hoon/KwebBlackBoard/blackboard-server/models"
	"hoon/KwebBlackBoard/blackboard-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBoard(c *gin.Context) {
	//req body로부터 course의 정보를 가져온다
	var body struct {
		Title string
		Desc  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})

		return
	}

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//token에서 user의 정보를 가져오고, 교수인지 확인한다(Isprofessor ==1)
	professorid := user.ID

	if !user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a professor!",
		})

		return
	}

	//parameter가 string이므로 uint로 변환하는 과정을 거친다
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course ID",
		})
		return
	}

	//교수가 가르치는 과목인지 확인한다
	var course models.Course
	result1 := initializers.DB.Joins("JOIN teaches ON teaches.course_id = courses.id").
		Where("teaches.professor_id = ? AND teaches.course_id = ?", professorid, courseid).
		First(&course)

	//교수가 가르치는 과목이 아니라면 에러를 출력한다
	if result1.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetching course",
		})

		return
	}

	if result1.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You don't teach this course",
		})

		return
	}

	//가져온 데이터를 이용해 새 board를 생성한다
	board := models.Board{CourseID: courseid, Title: body.Title, Desc: body.Desc}
	result2 := initializers.DB.Create(&board)

	if result2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create board",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created board",
	})

}

func SearchAllBoard(c *gin.Context) {
	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//token에서 user의 정보를 가져오고, 학생인지 확인한다(Isprofessor != true)
	studentid := user.ID

	if user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a student!",
		})

		return
	}

	//가져온 데이터를 이용해 조건에 맞는 boards를 불러온다
	var boards []models.Board
	result := initializers.DB.Joins("JOIN courses ON courses.id = boards.course_id").
		Joins("JOIN applies ON applies.course_id = courses.id AND applies.student_id =?", studentid).
		Where("applies.deleted_at IS NULL").
		Order("created_at DESC").
		Find(&boards)

	//에러가 있다면 에러를 출력한다
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetching boards",
		})

		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There are no board",
		})

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
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course ID",
		})
		return
	}

	//token에서 user의 정보를 가져오고, 학생인지 확인한다(Isprofessor != true)
	user_id := user.ID

	//교수일 경우, teach table 이용, 학생일 경우 apply table 이용
	if user.Isprofessor {
		var boards []models.Board
		result := initializers.DB.Joins("JOIN courses ON courses.id = boards.course_id").
			Joins("JOIN teaches ON teaches.course_id = courses.id AND teaches.professor_id =?", user_id).
			Where("teaches.course_id = ? AND teaches.deleted_at IS NULL", courseid).
			Find(&boards)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetching boards",
			})

			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "There are no board",
			})

			return
		}

		//Respond with them
		c.JSON(http.StatusOK, gin.H{
			"boards": boards,
		})
	} else { //학생일 경우, apply table 이용
		var boards []models.Board
		result := initializers.DB.Joins("JOIN courses ON courses.id = boards.course_id").
			Joins("JOIN applies ON applies.course_id = courses.id AND applies.student_id =?", user_id).
			Where("applies.course_id = ? AND applies.deleted_at IS NULL", courseid).
			Find(&boards)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetching boards",
			})

			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "There are no board",
			})

			return
		}

		//Respond with them
		c.JSON(http.StatusOK, gin.H{
			"boards": boards,
		})
	}
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

	//token에서 user의 정보를 가져오고, 학생인지 확인한다(Isprofessor != true)
	user_id := user.ID

	//교수일 경우, teach table 이용, 학생일 경우 apply table 이용
	if user.Isprofessor {
		var board []models.Board
		result := initializers.DB.Joins("JOIN courses ON courses.id = boards.course_id").
			Joins("JOIN teaches ON teaches.course_id = courses.id AND teaches.professor_id =?", user_id).
			Where("teaches.deleted_at IS NULL").
			First(&board, boardid)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetching boards",
			})

			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "There are no board or You're not relevant to this course",
			})

			return
		}

		//Respond with them
		c.JSON(http.StatusOK, gin.H{
			"board": board,
		})
	} else {
		var board []models.Board
		result := initializers.DB.Joins("JOIN courses ON courses.id = boards.course_id").
			Joins("JOIN applies ON applies.course_id = courses.id AND applies.student_id =?", user_id).
			Where("applies.deleted_at IS NULL").
			First(&board, boardid)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetching boards",
			})

			return
		}

		if result.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "There are no board or You're not relevant to this course",
			})

			return
		}

		//Respond with them
		c.JSON(http.StatusOK, gin.H{
			"board": board,
		})
	}
}
