package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wyoassign/wyoassign"
	"wyoassign/wyoclass"
)


func main() {
	wyoassign.InitAssignments()
	wyoclass.InitClasses()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.DeleteAssignment).Methods("DELETE")		
	router.HandleFunc("/assignment", wyoassign.CreateAssignment).Methods("POST")	
	router.HandleFunc("/assignments/{id}", wyoassign.UpdateAssignment).Methods("PUT")

	router.HandleFunc("/class-api-status", wyoclass.APISTATUS).Methods("GET")
	router.HandleFunc("/classes", wyoclass.GetClasses).Methods("GET")
	router.HandleFunc("/class/{id}", wyoclass.GetClass).Methods("GET")
	router.HandleFunc("/class", wyoclass.AddClass).Methods("POST")
	router.HandleFunc("/class/{id}", wyoclass.DropClass).Methods("DELETE")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}