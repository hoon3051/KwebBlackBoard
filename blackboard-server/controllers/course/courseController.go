package course

import (
	"net/http"

	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/utils"

	"github.com/gin-gonic/gin"
)

// 강의 생성 (Course와 Teach 생성)
func CreateCourse(c *gin.Context) {
	//req body로부터 course의 정보를 가져온다
	var body struct {
		Coursename   string
		Coursenumber string
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

	//가져온 데이터를 이용해 새 course를 생성한다
	course := models.Course{Coursename: body.Coursename, Coursenumber: body.Coursenumber}
	result := initializers.DB.Create(&course)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create a course",
		})

		return
	}

	//가져온 데이터를 이용해 새 teach를 생성한다

	courseid := course.ID

	teach := models.Teach{CourseID: courseid, ProfessorID: professorid}

	result2 := initializers.DB.Create(&teach)

	if result2.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create teach",
		})

		return
	}

	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Course created successfully",
	})
}

// 모든 강의 조회
func SearchAllCourse(c *gin.Context) {
	//강의들을 가져온다
	var courses []models.Course
	result := initializers.DB.Find(&courses)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There are no courses",
		})

		return
	}

	//보여준다
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}

// 등록(교수) 강의 조희
func SearchTeachCourse(c *gin.Context) {

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//token에서 user의 정보를 가져오고, 교수인지 확인한다(Isprofessor ==1)
	if !user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a professor!",
		})

		return
	}

	//Get the posts
	var courses []models.Course
	result := initializers.DB.Joins("JOIN teaches ON teaches.course_id = courses.id").Where("teaches.professor_id = ?", user.ID).Find(&courses)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetching courses",
		})

		return
	}

	if len(courses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You don't teach any course",
		})

		return
	}

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}

// 수강(학생) 강의 조희
func SearchApplyCourse(c *gin.Context) {

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	if user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a student!",
		})
	}

	//Get the posts
	var courses []models.Course
	result := initializers.DB.Joins("JOIN applies ON applies.course_id = courses.id").Where("applies.student_id = ?", user.ID).Find(&courses)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetching courses",
		})

		return
	}

	if len(courses) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "You don't apply any course",
		})

		return
	}

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}
