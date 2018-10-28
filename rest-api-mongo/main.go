package main

import(
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/google/uuid"
  "mgo"
  "dao"
)

//Create tasks structure
type Tasks struct{
  ID string `json:"id,omitempty"`
  Name  string `json:"taskName,omitempty"`
  Content string `json:"taskContent,omitempty"`
  State string `json:"taskState,omitempty"`
}

// create var with tasks structure
var list = []Tasks{}

// Functions
func GetList(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(list)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, task := range list {
    if task.ID == params["id"]{
      json.NewEncoder(w).Encode(task)
      return
    }
  }
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
  var task Tasks
  json.NewDecoder(r.Body).Decode(&task)
  uuid, _ := uuid.NewRandom()
  task.ID=fmt.Sprintf("%s",uuid)
  list = append(list, task)
  json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for id, task := range list {
    if task.ID == params["id"]{
      list = append(list[:id], list[id+1:]...)
    }
  }
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var task Tasks
  for i := 0; i < len(list); i++ {
    if list[i].ID == params["id"] {
      //parsea contenido de body a la variable task
      json.NewDecoder(r.Body).Decode(&task)
      list[i].State = task.State
    }
  }
}


// main function
func main(){
//add records to List slice
  router := mux.NewRouter()
  router.HandleFunc("/tasks", GetList).Methods("GET")
  router.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
  router.HandleFunc("/tasks", CreateTask).Methods("POST")
  router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")
  router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
  log.Fatal(http.ListenAndServe(":8000", router))
}
