package auth

import (
	"hoon/KwebBlackBoard/blackboard-server/initializers"
	"hoon/KwebBlackBoard/blackboard-server/models"

	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"golang.org/x/crypto/bcrypt"
)

// 로그인
func SignIn(c *gin.Context) {
	//req body로부터 username, password를 받아온다
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})

		return
	}

	//user의 정보를 찾아온다
	var user models.User
	initializers.DB.First(&user, "Username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username",
		})

		return
	}

	//password를 hashed password와 비교한다
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})

		return
	}

	//jwt token을 생성한다
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token",
		})

		return
	}

	//token을 보내준다
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully signed in",
	})

}

// 회원가입
func SignUp(c *gin.Context) {
	//req body로 부터 username, password, displayname, studentnumber, isprofessor를 가져온다
	var body struct {
		Username      string
		Password      string
		Displayname   string
		Studentnumber int32
		Isprofessor   bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body",
		})

		return
	}

	//password를 hash한다
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	//user를 body의 정보대로 생성한다
	user := models.User{Username: body.Username, Password: string(hash), Displayname: body.Displayname, Studentnumber: body.Studentnumber, Isprofessor: body.Isprofessor}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	//응답 보내준다
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created user",
	})
}

// 토큰이 있는지 확인
func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"Message": user,
	})
}
