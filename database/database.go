package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	// Customize it according to your database configuration
	// "{userDB}:{passwordDB}@tcp({hostDB}/{portDB})/{DBname}?parseTime=true"
	// parsetime to format times and dates with time.Time
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/todo_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
