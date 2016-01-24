package main

import (
  "encoding/json"
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  //fmt.Println("Response received at Index")
  fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  //fmt.Println("Response received at TodoIndex")
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Println("Response received at TodoShow:" + todoId)
  fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r * http.Request) {
  var todo todo

}
