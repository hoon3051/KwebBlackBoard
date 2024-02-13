package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/forms"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"

	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"errors"
)

type UserService struct{}

func (svc UserService) Login(c *gin.Context, loginForm forms.LoginForm) (user models.User, err error) {
	//user의 정보를 찾아온다
	initializers.DB.First(&user, "Username = ?", loginForm.Username)
	if user.ID == 0 {
		return user, errors.New("올바르지 않은 username입니다")
	}

	//password를 hashed password와 비교한다
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password))

	if err != nil {
		return user, errors.New("올바르지 않은 password입니다")
	}

	//jwt token을 생성한다
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return user, errors.New("failed to generate token")
	}

	//token을 보내준다
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	return user, nil
}

func (svc UserService) Register(registerForm forms.RegisterForm) (user models.User, err error) {
	//username이 이미 존재하는지 확인한다
	initializers.DB.First(&user, "Username = ?", registerForm.Username)
	if user.ID != 0 {
		return user, errors.New("이미 존재하는 username입니다")
	}

	//password를 hash한다
	hash, err := bcrypt.GenerateFromPassword([]byte(registerForm.Password), 10)

	if err != nil {
		return user, errors.New("failed to hash password")
	}

	//user를 body의 정보대로 생성한다
	user = models.User{
		Username:      registerForm.Username,
		Password:      string(hash),
		Displayname:   registerForm.Displayname,
		Studentnumber: registerForm.Studentnumber,
		Isprofessor:   registerForm.Isprofessor,
	}

	//user를 db에 저장한다
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		return user, errors.New("가입에 실패했습니다")
	}

	return user, nil

}
