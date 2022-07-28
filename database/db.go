package database

import (
	"fmt"
	"go-chat/handlers"
	"go-chat/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	var config models.Config
	handlers.ReadFile(&config, "./resources/config.yml")
	dsn := "host=" + config.Database.Host +
		" user=" + config.Database.User +
		" password=" + config.Database.Password +
		" dbname=" + config.Database.DbName +
		" port=" + config.Database.Port +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func MigrateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.UserModel{}); err != nil {
		return err
	}
	return nil
}

func CreateTables(db *gorm.DB) error {
	if !db.Migrator().HasTable(&models.UserModel{}) {
		if err := db.Migrator().CreateTable(&models.UserModel{}); err != nil {
			return err
		}
		fmt.Println("Create table success")
	}
	// seeder
	var users []models.UserModel = []models.UserModel{
		{Name: "admin", Email: "admin@gmail.com"},
		{Name: "user1", Email: "user1@gmail.com"},
	}
	for _, user := range users {
		db.Create(&user)
		fmt.Println("Create user success: ", user.Name)
	}
	return nil
}
