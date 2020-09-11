package models

import (
	"log"
	"os"

	"regisapp/controllers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GetConnectionDB return the database connection
func GetConnectionDB() *gorm.DB {
	dsn := "user=postgres password=admin dbname=regisapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR CONNECTING TO DATABASE : %v", err.Error())
		os.Exit(100)
	}
	log.Printf("CONNECTED TO DATABASE")

	controllers.CreateuserTable(db)
	controllers.InitiateDB(db)

	return db
}

/*
other configurations
// https://github.com/go-gorm/postgres
db, err := gorm.Open(postgres.New(postgres.Config{
	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})
*/
