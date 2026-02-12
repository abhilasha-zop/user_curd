package main

import (
	"database/sql"
	"log"
	"net/http"
	"user_crud/handler"

	_ "github.com/go-sql-driver/mysql"
	//"user_crud/model"
	"user_crud/service"
	"user_crud/store"
)

func main() {
	db, _ := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")

	store := &store.UserStore{DB: db}
	service := &service.UserService{Store: store}
	handler := &handler.UserHandler{Service: service}

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/users", handler.GetAll)
	http.HandleFunc("/user/", handler.GetOne)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete/", handler.Delete)

	log.Println("server :8080")
	http.ListenAndServe(":8080", nil)
}
