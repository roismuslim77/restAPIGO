package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/user",returnUsers).Methods("GET")
	router.HandleFunc("/user/new", insertUser).Methods("POST")
	router.HandleFunc("/user", showUser).Methods("POST")
	router.HandleFunc("/user", editUser).Methods("PUT")
	router.HandleFunc("/user", delUser).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Conect port 1234")
	log.Fatal(http.ListenAndServe(":1234",router))
}
