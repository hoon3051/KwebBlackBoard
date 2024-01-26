package routers

import (
	"time"
	
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/controllers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/middlewares"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middlewares.RequireAuth, controllers.Validate)
}

func CourseRoutes(router *gin.Engine) {
	router.Use(middlewares.RequireAuth)
	router.POST("/course", controllers.CreateCourse)
	router.GET("/course", controllers.SearchAllCourse)
	router.GET("/course/professor", controllers.SearchTeachCourse)
	router.GET("/course/student", controllers.SearchApplyCourse)
}

func ApplyRoutes(router *gin.Engine) {
	router.Use(middlewares.RequireAuth)
	router.POST("/apply/:course_id", controllers.ApplyCourse)
	router.GET("/apply/:course_id", controllers.SearchAppliedStudent)
	router.DELETE("/apply/:course_id", controllers.DeleteAppliedStudent)
}

func BoardRoutes(router *gin.Engine) {
	router.Use(middlewares.RequireAuth)
	router.POST("/board/:course_id", controllers.CreateBoard)
	router.GET("/board/", controllers.SearchAllBoard)
	router.GET("/board/:course_id", controllers.SearchCourseBoard)
	router.GET("/board/course/:board_id", controllers.GetBoard)
}

func RouterSetupV1() *gin.Engine {
	r := gin.Default()


	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 클라이언트의 도메인 명시
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true, // withCredentials 요청 허용
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	UserRoutes(r)
	CourseRoutes(r)
	ApplyRoutes(r)
	BoardRoutes(r)

	return r
}