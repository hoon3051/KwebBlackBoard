package apply

import (
	"hoon/KwebBlackBoard/blackboard-server/initializers"
	"hoon/KwebBlackBoard/blackboard-server/models"
	"hoon/KwebBlackBoard/blackboard-server/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyCourse(c *gin.Context) {

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//token에서 user의 정보를 가져오고, 학생인지 확인한다(Isprofessor != false)
	studentid := user.ID

	if user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a student!",
		})

		return
	}

	//가져온 데이터를 이용해 새 teach를 생성한다

	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course ID",
		})
		return
	}

	//가져온 데이터를 이용해 새 apply를 생성한다
	apply := models.Apply{CourseID: courseid, StudentID: studentid}

	result := initializers.DB.Create(&apply)

	//에러가 있다면 에러를 출력한다
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create apply",
		})

		return
	}

	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Apply created successfully",
	})
}

func SearchAppliedStudent(c *gin.Context) {
	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//token에서 user의 정보를 가져오고, 교수인지 확인한다(Isprofessor == true)
	professorid := user.ID

	if !user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a professor!",
		})

		return
	}

	//가져온 데이터를 이용해 강의(course)가 자신이 가르치는(teach) 강의인지 확인한다

	//parameter가 string이므로 uint로 변환하는 과정을 거친다
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course ID",
		})
		return
	}

	//가져온 데이터를 이용해 새 course를 생성한다
	var course models.Course
	result1 := initializers.DB.Joins("JOIN teaches ON teaches.course_id = courses.id").
		Where("teaches.professor_id = ? AND teaches.course_id = ?", professorid, courseid).
		First(&course)

	//에러가 있다면 에러를 출력한다
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

	//apply를 이용해 해당 course의 학생들(user)을 가져온다
	var students []models.User
	result2 := initializers.DB.Joins("JOIN applies ON users.id = applies.student_id").
		Where("applies.course_id = ? AND applies.deleted_at IS NULL", courseid).
		Find(&students)

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetching users",
		})

		return
	}

	if result2.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There are no student that apply this course",
		})

		return
	}

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"users": students,
	})
}

func DeleteAppliedStudent(c *gin.Context) {

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//token에서 user의 정보를 가져오고, 교수인지 확인한다(Isprofessor == true)
	professorid := user.ID

	if !user.Isprofessor {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User is not a professor!",
		})

		return
	}

	// 가져온 데이터를 이용해 강의(course)가 자신이 가르치는(teach) 강의인지 확인한다

	// parameter가 string이므로 uint로 변환하는 과정을 거친다
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course ID",
		})
		return
	}

	// 가져온 데이터를 이용해 새 course를 생성한다
	var course models.Course
	result1 := initializers.DB.Joins("JOIN teaches ON teaches.course_id = courses.id").
		Where("teaches.professor_id = ? AND teaches.course_id = ?", professorid, courseid).
		First(&course)


	// 에러가 있다면 에러를 출력한다
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

	var body struct {
		Studentid uint
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})

		return
	}

	// apply를 이용해 해당 course의 학생들(user)을 가져온다
	result2 := initializers.DB.Where("course_id = ? AND student_id = ?", courseid, body.Studentid).Delete(&models.Apply{})

	// 에러가 있다면 에러를 출력한다
	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete applied student",
		})
	}

	if result2.RowsAffected == 0 {
		log.Printf("Error fetching applies: %v\n", result2.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No applied student found or already deleted",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Applied student deleted successfully",
	})
}
