package api

import (
  "encoding/json"
  "log"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/golang-explore/mongo-go-api/db"
  )



var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request)  {
  json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request)  {
  params :=mux.Vars(req)
  for _, item := range people {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request)  {
  params := mux.Vars(req)
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = params["id"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request)  {
  params := mux.Vars(req)
  for index, item := range people {
    if item.ID == params["id"] {
      people = append(people[:index], people[index+1:]...)
      break
      }
  }
  json.NewEncoder(w).Encode(people)
}
