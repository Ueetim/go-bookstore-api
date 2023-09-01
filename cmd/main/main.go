package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Ueetim/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	log.Println("Starting server on :9010")
	err := http.ListenAndServe(":9010", r)
	log.Fatal(err)
}