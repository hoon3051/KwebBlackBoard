package controllers

import (
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/forms"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/services"

	"net/http"

	"github.com/gin-gonic/gin"
	
)

// 로그인
func Login(c *gin.Context) {
	//req body로부터 username, password를 받아온다 (form)
	var loginForm forms.LoginForm
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//validation을 한다 (form)
	userForm := forms.UserForm{}
	if err := userForm.Login(loginForm); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
		return
	}

	//login한다 (service)
	userService := services.UserService{}
	user, err := userService.Login(c, loginForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	
	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed in",
		"user": user,
	})

}

// 회원가입
func Register(c *gin.Context) {
	//req body로 부터 username, password, displayname, studentnumber, isprofessor를 가져온다 (form)
	var registerForm forms.RegisterForm
	if err := c.ShouldBindJSON(&registerForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//validation을 한다 (form)
	userForm := forms.UserForm{}
	if err := userForm.Register(registerForm); err != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
		return
	}

	//register한다 (service)
	userService := services.UserService{}
	user, err := userService.Register(registerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}


	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created user",
		"user": user,
	})
}

// 토큰이 있는지 확인
func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"Message": user,
	})
}
