package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/Foodut/backend/controllers"
	model "github.com/Foodut/backend/models"
)

func main() {
	db := controller.Connect()
	db.Debug().AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Customer{},
		&model.Seller{},
		&model.DetailProduct{},
		&model.Transaction{},
		&model.Category{},
		&model.Product{},
		&model.Picture{},
	)
	router := mux.NewRouter()

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
