package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select * from employee")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.Name, &users.City); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_user = append(arr_user, users)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func insertUser(w http.ResponseWriter, r *http.Request)  {
	var response Response

	db := connect()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	firstName := r.FormValue("name")
	city := r.FormValue("city")

	inForm, err := db.Prepare("INSERT INTO `employee`(`id`, `name`, `city`) VALUES (null,?,?)")
	if err != nil {
		panic(err)
	}
	inForm.Exec(firstName,city)

	response.Status = 201
	response.Message = "succes add"
	log.Print("Insert to db")

	defer db.Close()
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func showUser(w http.ResponseWriter, r *http.Request)  {
	var response Response
	user := Users{}
	var arr_user []Users

	db := connect()
	id := r.URL.Query().Get("id")
	shQuery, err := db.Query("SELECT * FROM `employee` where `id` = ?", id)
	if err != nil {
		panic(err)
	}

	for shQuery.Next() {
		var id, name, city string
		if err := shQuery.Scan(&id, &name, &city); err != nil {
			log.Fatal(err.Error())
		}
		user.Id = id
		user.Name = name
		user.City = city
	}
	arr_user = append(arr_user, user)
	response.Status = 200
	response.Message = "success show"
	response.Data = arr_user

	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func editUser(w http.ResponseWriter, r *http.Request)  {
	var response Response

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	
	db := connect()
	id := r.URL.Query().Get("id")
	name := r.FormValue("name")
	city := r.FormValue("city")

	query, err := db.Prepare("UPDATE `employee` SET `name`= ?, `city`= ? WHERE `id`= ?")
	if err != nil {
		panic(err)
	} 
	query.Exec(name, city, id)
	response.Status = 201
	response.Message = "update db"
	log.Print("updated to db "+name+" "+city)

	defer db.Close()
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func delUser(w http.ResponseWriter, r *http.Request) {
	db:= connect()
	var response Response

	id := r.URL.Query().Get("id")
	que, err := db.Prepare("DELETE FROM `employee` WHERE `id`= ?")
	if err != nil {
		panic(err)
	}
	
	que.Exec(id)
	response.Status = 200
	response.Message = "Delete succes"
	log.Print("user delete")

	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}