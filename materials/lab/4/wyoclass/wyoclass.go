package wyoclass

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Response struct {
	Classes []Course `json:"classes"`
}

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"desc"`
	Credits int `json:"credits"`
	Grade string `json:"grade"`
}

var Classes []Course

func InitClasses(){
	var course Course
	course.Id = "4010"
	course.Name = "Special Topics: Cybersecurity"
	course.Description = "Black Hat Go"
	course.Credits = 3
	course.Grade = "A"
	Classes = append(Classes, course)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}


func GetClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, class := range Classes {
		if class.Id == params["id"] {
			json.NewEncoder(w).Encode(class)
			break
		}
	}
}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	
	var response Response

	response.Classes = Classes

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func AddClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var course Course
	r.ParseForm()

	if r.FormValue("id") != "" {
		course.Id = r.FormValue("id")
		course.Name = r.FormValue("name")
		course.Description = r.FormValue("desc")
		course.Credits, _ = strconv.Atoi(r.FormValue("credits"))
		course.Grade = r.FormValue("grade")
		Classes = append(Classes, course)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)
}

func DropClass(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"

	for index, course := range Classes {
		if course.Id == params["id"] {
			Classes = append(Classes[:index], Classes[index+1:]...)
			response["status"] = "Success"
			break
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}