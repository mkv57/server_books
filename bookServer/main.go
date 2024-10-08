package main

import (
	"bookServer/internal/api"
	"bookServer/internal/db"
	"bookServer/internal/domain"

	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	file, err := os.OpenFile("app.log", os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	r := mux.NewRouter()

	r.Use(api.Logging(api.Logger))

	ourServer := api.Server{
		Database: db.Repository{
			Store: make(map[int]domain.Book),
		},
	}

	r.HandleFunc("/book", ourServer.GetBook).Methods(http.MethodGet)
	r.HandleFunc("/book", ourServer.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/book", ourServer.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/book", ourServer.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books", ourServer.AllBooks).Methods(http.MethodGet)

	api.Logger.Warn("сервер запущен")
	fmt.Println("сервер запущен")
	err = http.ListenAndServe("127.0.0.1:8080", r)
	api.Logger.Warn("сервер отключён")
	if err != nil {
		api.Logger.Error("сервер нe запустился")
		log.Fatal(err)
	}

}
