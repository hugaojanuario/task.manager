package database

import (
	"log"

	"github.com/hugaojanuario/task.manager.api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
	err error
)

func ConectingOnDatabase(){
	stringConection := "host=localhost user=admin password=123456 dbname=taskmanager port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConection))
	if err != nil{
		log.Panic("ERROR: ", err)
	}

	DB.AutoMigrate(&model.Tasks)
	
}