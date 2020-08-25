package controller

import (
	"../database"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
)

const RootURL = "http://localhost:8000/"

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("./view/index.html")
	anime := database.AllAnime()
	t.Execute(w, anime)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	database.CreateAnime(title, description)
	http.Redirect(w, r, RootURL, 302)
}

func Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id_s := ps.ByName("id")
	id, _ := strconv.Atoi(id_s)
	anime := database.FindAnimeById(id)
	t, _ := template.ParseFiles("./view/edit.html")
	t.Execute(w, anime)
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	id, _ := strconv.Atoi(ps.ByName("id"))
	database.UpdateAnime(id, title, description)
	http.Redirect(w, r, RootURL, 302)
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	database.DeleteAnime(id)
	http.Redirect(w, r, RootURL, 302)
}
