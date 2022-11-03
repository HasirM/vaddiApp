package database

import (
	"fmt"
	"vaddi/models"

	"github.com/hasirm/vaddi/controllers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := controllers.GetENV["USERNAME"] + ":" + controllers.GetENV["PASSWORD"] + "@/" + controllers.GetENV["DB_NAME"]

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(dsn)
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
