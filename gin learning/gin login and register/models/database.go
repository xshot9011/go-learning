package models

import (
	"log"
	"os"
	"time"

	"regisapp/controllers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User model for db
type User struct {
	UUID         string    `gorm:"primaryKey"`
	Fullname     string    `gorm:"not null;size:256"`
	Email        string    `gorm:"not null;unique;size:256"`
	PasswordHash string    `gorm:"not null;size:256"`
	Salt         string    `gorm:"not null;size:32"`
	Created      time.Time `gorm:"autoCreateTime:milli"`
	Updated      time.Time `gorm:"autoUpdateTime:milli"`
}

// GetConnectionDB return the database connection
func GetConnectionDB() *gorm.DB {
	dsn := "user=postgres password=admin dbname=regisapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR CONNECTING TO DATABASE : %v", err.Error())
		os.Exit(100)
	}
	log.Printf("CONNECTED TO DATABASE")

	createTable(db)
	controllers.InitiateDB(db)

	return db
}

func createTable(db *gorm.DB) {
	isExist := db.Migrator().HasTable(&User{})
	if !isExist {
		db.Migrator().CreateTable(&User{})
	}
}

/*
other configurations
// https://github.com/go-gorm/postgres
db, err := gorm.Open(postgres.New(postgres.Config{
	DSN: "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
	PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})
*/
