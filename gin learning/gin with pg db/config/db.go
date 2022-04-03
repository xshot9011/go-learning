package config

import (
	"log"
	"os"

	controllers "github.com/cavdy-play/go_db/controllers"
	"github.com/go-pg/pg/v9"
)

// GetConnetion return the connection to db
func GetConnetion() *pg.DB {
	options := &pg.Options{
		User:     "postgres",
		Password: "admin",
		Addr:     "localhost:5432",
		Database: "todo",
	}

	var db *pg.DB = pg.Connect(options)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	// when database is connected, then create the table
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)

	return db
}
