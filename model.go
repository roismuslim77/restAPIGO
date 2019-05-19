package main

type Users struct {
	Id string `form:"id" json:"id"`
	Name string `form:"name" json:"nama"`
	City string `form:"city" json:"kota"`
}

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data []Users
}