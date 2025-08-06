package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	var user *person = &person{
		Name: "Sahabat Devetek",
		Age:  15,
	}

	w.Header().Set("content-type", "application/json")
	j, _ := json.Marshal(user)
	w.Write(j)
}

func main() {
	http.HandleFunc("/", httpHandler)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
