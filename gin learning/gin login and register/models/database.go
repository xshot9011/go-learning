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
	UUID            string    `gorm:"primaryKey" form:"-"`
	Fullname        string    `gorm:"not null;size:256" form:"fullname" validate:"required,min=1,max=256"`
	Email           string    `gorm:"not null;unique;size:256" form:"email" validate:"required,min=4,max=256,email"`
	Password        string    `gorm:"-" form:"password" validate:"required,min=8,eqfield=PasswordConfirm"`
	PasswordConfirm string    `gorm:"-" form:"password_confirm" validate:"required,min=8"`
	PasswordHash    string    `gorm:"not null;size:256" form:"-"`
	Salt            string    `gorm:"not null;size:32" form:"-"`
	Created         time.Time `gorm:"autoCreateTime:milli" form:"-"`
	Updated         time.Time `gorm:"autoUpdateTime:milli" form:"-"`
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
