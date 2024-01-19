package controllers

import (
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/controllers/apply"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/controllers/auth"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/controllers/board"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/controllers/course"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/sign_up", auth.SignUp)
	router.POST("/sign_in", auth.SignIn)
	router.GET("/validate", middlewares.RequireAuth, auth.Validate)
}

func CourseRoutes(router *gin.Engine) {
	router.POST("/course", middlewares.RequireAuth, course.CreateCourse)
	router.GET("/course", course.SearchAllCourse)
	router.GET("/course/professor", middlewares.RequireAuth, course.SearchTeachCourse)
	router.GET("/course/student", middlewares.RequireAuth, course.SearchApplyCourse)
}

func ApplyRoutes(router *gin.Engine) {
	router.POST("/apply/:course_id", middlewares.RequireAuth, apply.ApplyCourse)
	router.GET("/apply/:course_id", middlewares.RequireAuth, apply.SearchAppliedStudent)
	router.DELETE("/apply/:course_id", middlewares.RequireAuth, apply.DeleteAppliedStudent)
}

func BoardRoutes(router *gin.Engine) {
	router.POST("/board/:course_id", middlewares.RequireAuth, board.CreateBoard)
	router.GET("/board/", middlewares.RequireAuth, board.SearchAllBoard)
	router.GET("/board/:course_id", middlewares.RequireAuth, board.SearchCourseBoard)
	router.GET("/board/course/:board_id", middlewares.RequireAuth, board.GetBoard)
}
