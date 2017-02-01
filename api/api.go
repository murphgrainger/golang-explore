package api

import (
  "encoding/json"
  "log"
  "net/http"

  "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"


  "github.com/gorilla/mux"
  "github.com/murphgrainger/goproject/db"
  "strconv"

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

// func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request)  {
//   json.NewEncoder(w).Encode(people)
// }
//
// func GetPersonEndpoint(w http.ResponseWriter, req *http.Request)  {
//   params :=mux.Vars(req)
//   for _, item := range people {
//     if item.ID == params["id"] {
//       json.NewEncoder(w).Encode(item)
//       return
//     }
//   }
//   json.NewEncoder(w).Encode(&Person{})
// }
//
// func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request)  {
//   params := mux.Vars(req)
//   var person Person
//   _ = json.NewDecoder(req.Body).Decode(&person)
//   person.ID = params["id"]
//   people = append(people, person)
//   json.NewEncoder(w).Encode(people)
// }
//
// func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request)  {
//   params := mux.Vars(req)
//   for index, item := range people {
//     if item.ID == params["id"] {
//       people = append(people[:index], people[index+1:]...)
//       break
//       }
//   }
//   json.NewEncoder(w).Encode(people)
// }



func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}
