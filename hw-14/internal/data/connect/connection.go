package connect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DataBaseConnect() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:33306)/tours")
	if err != nil {
		fmt.Println("Error opening database")
		return nil
	}

	return db
}
