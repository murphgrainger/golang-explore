package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/murphgrainger/goproject/api"
  )

func main()  {
  router := mux.NewRouter()
  router.HandleFunc("/api/people", api.GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/api/people/{id}", api.GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/api/people", api.CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/api/people/{id}", api.DeletePersonEndpoint).Methods("DELETE")
  http.ListenAndServe(":3000", router)
}
