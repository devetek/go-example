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
	var nakama *person = &person{
		Name: "Nakama",
		Age:  30,
	}

	w.Header().Set("content-type", "application/json")
	j, _ := json.Marshal(nakama)
	w.Write(j)
}

func main() {
	http.HandleFunc("/", httpHandler)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
