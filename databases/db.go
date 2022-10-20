package databases

import (
	"assignment-2/entity"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "12345"
	dbPort   = 5432
	dbName   = "order_by"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		fmt.Println("error open connection to database", err.Error())
	}

	err = db.Debug().AutoMigrate(entity.Order{}, entity.Item{})

	if err != nil {
		fmt.Println("error on migration", err.Error())
	}

	fmt.Println("successfully connected to my database")

}

func GetDB() *gorm.DB {
	return db
}
