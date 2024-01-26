package main

import (
	"fmt"

	"github.com/hoon3051/KwebBlackBoard/blackboard-server/initializers"
	"github.com/hoon3051/KwebBlackBoard/blackboard-server/routers"

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

	router := routers.RouterSetupV1()


	router.Run()
}
