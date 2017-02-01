package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/murphgrainger/goproject/api"
  )

func main()  {
  router := mux.NewRouter()
  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  // router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
  // router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
  // router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
  http.ListenAndServe(":3000", router)
}
