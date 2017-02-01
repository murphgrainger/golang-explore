package main

import (
  "encoding/json"
  "log"
  "net/http"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/murphgrainger/goproject/api"
  )

func main()  {
  session, err := mgo.Dial(url)
  c := session.DB(database).C(collection)
  err := c.Find(query).One(&result)
  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":3000", router))
}
