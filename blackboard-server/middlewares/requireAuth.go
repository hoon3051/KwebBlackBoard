package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	//req로부터 cookie를 가져온다
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//cookie를 decode 및 validate한다
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//exp를 확인한다
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//token의 sub로 user를 찾는다
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//req에 첨부한다
		c.Set("user", user)

		//다음 middleware로 넘긴다
		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
