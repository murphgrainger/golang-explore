package api

import (
  "net/http"

  "encoding/json"
  "fmt"
  "github.com/murphgrainger/goproject/db"
  "github.com/gorilla/mux"

  )
  // "strconv"




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

// func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
// 	ID := req.FormValue("id")
// 	valueStr := req.FormValue("value")
// 	value, err := strconv.Atoi(valueStr)
// 	if err != nil {
// 		handleError(err, "Failed to parse input data: %v", w)
// 		return
// 	}
//
// 	item := db.Item{ID: ID, Value: value}
//
// 	if err = db.Save(item); err != nil {
// 		handleError(err, "Failed to save data: %v", w)
// 		return
// 	}
//
// 	w.Write([]byte("OK"))
// }

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
