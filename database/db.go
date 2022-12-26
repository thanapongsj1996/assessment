package database

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB(connectionStr string) {
	var err error

	db, err = gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	CreateTableExpenses()
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	_db, _ := db.DB()
	_db.Close()
}

func CreateTableExpenses() {
	createTableScript := `CREATE TABLE IF NOT EXISTS expenses (id SERIAL PRIMARY KEY,title TEXT,amount FLOAT,note TEXT,tags TEXT[]);`
	tx := db.Exec(createTableScript)
	if err := tx.Error; err != nil {
		log.Fatal("can't create table", err)
	}
}
