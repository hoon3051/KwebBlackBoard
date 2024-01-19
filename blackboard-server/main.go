package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/controllers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"

	"github.com/gin-gonic/gin"
)

// initializers에 있는 함수들 실행
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

// main
func main() {
	fmt.Println("hello")

	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 클라이언트의 도메인 명시
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true, // withCredentials 요청 허용
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	controllers.AuthRoutes(router)
	controllers.CourseRoutes(router)
	controllers.ApplyRoutes(router)
	controllers.BoardRoutes(router)

	router.Run()
}
