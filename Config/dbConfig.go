package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(){
	var err error
	dsn := "root:root@tcp(127.0.0.0:3306)/hospify"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected Successfully")
}
