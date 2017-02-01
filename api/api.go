package api

import (
  "net/http"

  "encoding/json"
  "fmt"
  "github.com/murphgrainger/goproject/db"
  "github.com/gorilla/mux"

  )


  // GetAllItems returns a list of all database items to the response.
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAll()
	if err != nil {
		handleError(err, "Failed to load database items: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOne(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
  Firstname := req.FormValue("firstname")
  Lastname := req.FormValue("lastname")
  City := req.FormValue("city")


  person:= db.Person{ID: ID, Firstname: Firstname, Lastname: Lastname, City: City}

  if err = db.Save(person); err != nil {
  		handleError(err, "Failed to save data: %v", w)
  		return
  	}

	w.Write([]byte("OK"))
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.Remove(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}


func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}
