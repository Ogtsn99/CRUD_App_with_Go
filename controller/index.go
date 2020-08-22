package controller

import (
	"../database"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

const RootURL = "http://localhost:8000/"

type Anime struct {
	ID int
	Title string
	Description string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	t, _ := template.ParseFiles("./view/index.html")
	rows, err := database.DB.Query("SELECT id, title, description FROM anime")

	if err != nil {
		panic(err)
	}

	var animes []Anime
	for rows.Next(){
		var id int
		var title string
		var description string
		rows.Scan(&id, &title, &description)
		newAnime := Anime{id, title, description}
		animes = append(animes, newAnime)
	}
	t.Execute(w, animes)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	sqlStatement := `
 	INSERT INTO anime (title, description)
 	VALUES ($1, $2)`
	_, err := database.DB.Query(sqlStatement, title, description)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, RootURL, 302)
}

func Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id_s := ps.ByName("id")
	var title string
	var description string
	sqlStatement := `
 	SELECT title, description FROM anime
 	WHERE id=$1`
	row := database.DB.QueryRow(sqlStatement, id_s)
	err := row.Scan(&title, &description)
	if err != nil {
		panic(err)
	}
	var id_i int
	id_i, _ = strconv.Atoi(id_s)
	anime := Anime{id_i, title, description}
	t, _ := template.ParseFiles("./view/edit.html")
	t.Execute(w, anime)
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	id := ps.ByName("id")
	sqlStatement := `
 	UPDATE anime
 	SET title = $2, description = $3
 	WHERE id = $1;`
	_, err := database.DB.Exec(sqlStatement, id, title, description)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, RootURL, 302)
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	sqlStatement := `
	DELETE FROM anime WHERE id = $1;`
	_, err := database.DB.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, RootURL, 302)
}
