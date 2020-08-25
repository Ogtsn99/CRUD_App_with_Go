package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

type Anime struct {
	ID int
	Title string
	Description string
}

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

func AllAnime()[]Anime {
	rows, err := DB.Query("SELECT id, title, description FROM anime")
	if err != nil {
		panic(err)
	}

	var anime []Anime
	for rows.Next(){
		var id int
		var title string
		var description string
		rows.Scan(&id, &title, &description)
		newAnime := Anime{id, title, description}
		anime = append(anime, newAnime)
	}
	return anime
}

func CreateAnime(title string, description string) {
	sqlStatement := `
 	INSERT INTO anime (title, description)
 	VALUES ($1, $2)`
	_, err := DB.Query(sqlStatement, title, description)
	if err != nil {
		panic(err)
	}
}

func FindAnimeById(id int) Anime{
	var title string
	var description string
	sqlStatement := `
 	SELECT title, description FROM anime
 	WHERE id=$1`
	row := DB.QueryRow(sqlStatement, id)
	err := row.Scan(&title, &description)
	if err != nil {
		panic(err)
	}
	return Anime{id, title, description}
}

func UpdateAnime(id int, title string, description string) {
	sqlStatement := `
 	UPDATE anime
 	SET title = $2, description = $3
 	WHERE id = $1;`
	_, err := DB.Exec(sqlStatement, id, title, description)
	if err != nil {
		panic(err)
	}
}

func DeleteAnime(id int){
	sqlStatement := `
	DELETE FROM anime WHERE id = $1;`
	_, err := DB.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
}
