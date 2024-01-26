package initializers

import (
	"log"

	"github.com/joho/godotenv"
)


//.env파일 main에 연결
func LoadEnvVariables() {
	err := godotenv.Load()

	if err !=nil{
		log.Fatal("Failed to load .env file")
	}
}