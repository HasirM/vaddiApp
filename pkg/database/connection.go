package database

import (
	"fmt"
	// "os"

	"github.com/hasirm/vaddiapp/pkg/models"
	"github.com/hasirm/vaddiapp/pkg/config"

	// "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.Data().DBUsername + ":" + config.Data().DBPassword + "@/" + config.Data().DBName
	// dsn := "root:password@/vaddi_db"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(dsn)
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
