package controllers

import (
	"net/http"

	"github.com/hoon3051/KwebBlackBoard/blackboard-server/forms"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/services"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/utils"

	"github.com/gin-gonic/gin"
)

// 강의 생성 (Course와 Teach 생성)
func CreateCourse(c *gin.Context) {
	//req body로부터 course의 정보를 가져온다
	var createCourseForm forms.CreateCourseForm
	if err := c.ShouldBindJSON(&createCourseForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validation을 한다 (form)
	courseFormm := forms.CourseForm{}
	if err := courseFormm.Create(createCourseForm); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//course를 생성한다 (service)
	courseService := services.CourseService{}
	course, err := courseService.CreateCourse(*user, createCourseForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Course created successfully",
		"course":  course,
	})
}

// 모든 강의 조회
func SearchAllCourse(c *gin.Context) {
	//강의들을 가져온다 (service)
	courseService := services.CourseService{}
	courses, err := courseService.SearchAllCourse()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	//course를 가져온다 (service)
	courseService := services.CourseService{}
	courses, err := courseService.SearchTeachCourse(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	//course를 가져온다 (service)
	courseService := services.CourseService{}
	courses, err := courseService.SearchApplyCourse(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Respond with them
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}
