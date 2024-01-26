package controllers

import (
	"net/http"

	"github.com/hoon3051/KwebBlackBoard/blackboard-server/utils"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/services"

	"github.com/gin-gonic/gin"
)

func ApplyCourse(c *gin.Context) {
	//token에서 user의 정보를 가져오고, 에러가 있다면 에러를 출력한다.
	user, statusCode, err := utils.GetUser(c)
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//param에서 courseID를 가져온다
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	//새 apply를 생성한다 (service)
	applyService := services.ApplyService{}
	apply, err := applyService.CreateApply(*user, courseid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Apply created successfully",
		"apply":   apply,
	})
}

func SearchAppliedStudent(c *gin.Context) {
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

	//강의(course)가 자신이 가르치는(teach) 강의인지 확인한다 (service)
	teachService := services.TeachService{}
	_, err = teachService.GetTeach(*user, courseid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//apply를 이용해 수강한 학생들의 정보를 가져온다 (service)
	applyService := services.ApplyService{}
	students, err := applyService.GetAppliedStudent(courseid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//응답해준다
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

	// param으로 courseID를 가져온다
	courseid, err := utils.GetUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid course ID",
		})
		return
	}

	// 가져온 데이터를 이용해 강의(course)가 자신이 가르치는(teach) 강의인지 확인한다
	teachService := services.TeachService{}
	_, err = teachService.GetTeach(*user, courseid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//body에서 studentID를 가져온다
	var body struct {
		Studentid uint
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the body"})
		return
	}

	//가져온 데이터를 이용해 apply를 삭제한다 (service)
	applyService := services.ApplyService{}
	err = applyService.DeleteApply(*user, body.Studentid, courseid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Applied student deleted successfully"})
}
