package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func Init(){
	host := "localhost"
	port := 5432
	//ユーザーネームを環境変数で指定
	user := os.Getenv("USERNAME")
	dbname := "animetasks"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "dbname=%s sslmode=disable", host, port, user, dbname)
	var err interface{}
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
}
