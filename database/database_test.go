package database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/todo_golang?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connected")
}
