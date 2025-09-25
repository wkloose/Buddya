package initializers

import (
	"fmt"
	"os"
	"github.com/wkloose/Buddya/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println("failed to connect to database")
	}
}

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.City{},
		&models.Place{},
		&models.PromptLog{},
		&models.AIEvent{},
		&models.PromptTemplate{},
		&models.Event{},
	)
	if err != nil {
		fmt.Println("failed to migrate database")
	}
}
func ResetDatabase() {
	DB.Migrator().DropTable(
		&models.User{},
		&models.City{},
		&models.Place{},
		&models.PromptLog{},
		&models.AIEvent{},
		&models.PromptTemplate{},
		&models.Event{},
	)

	SyncDatabase()
}