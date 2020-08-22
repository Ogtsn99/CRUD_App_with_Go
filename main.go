package main

import (
	"./controller"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"./database"
)

func main() {
	database.Init()
	router := httprouter.New()
	router.GET("/", controller.Index)
	router.POST("/", controller.Create)
	router.GET("/:id", controller.Edit)
	router.POST("/:id/update", controller.Update)
	router.POST("/:id/delete", controller.Delete)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
