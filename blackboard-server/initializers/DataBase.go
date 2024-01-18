package initializers

import (
	"fmt"
	"hoon/KwebBlackBoard/blackboard-server/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// .env에서 정의한 mysql database에 연결
func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to DB")
	}
}

// models DB에 연동(migrate)
func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Course{})
	DB.AutoMigrate(&models.Teach{})
	DB.AutoMigrate(&models.Apply{})
	DB.AutoMigrate(&models.Board{})
}
