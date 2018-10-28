package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/src/models"
)

// create var with tasks structure
var list = []models.Tasks{}

// Functions
func GetList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(list)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, task := range list {
		if task.ID == params["id"] {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Tasks
	json.NewDecoder(r.Body).Decode(&task)
	uuid, _ := uuid.NewRandom()
	task.ID = fmt.Sprintf("%s", uuid)
	list = append(list, task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for id, task := range list {
		if task.ID == params["id"] {
			list = append(list[:id], list[id+1:]...)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Tasks
	for i := 0; i < len(list); i++ {
		if list[i].ID == params["id"] {
			//parsea contenido de body a la variable task
			json.NewDecoder(r.Body).Decode(&task)
			list[i].State = task.State
		}
	}
}

// main function
func main() {
	//add records to List slice
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetList).Methods("GET")
	router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
