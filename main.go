package main

import (
	"bitbucket.org/new-joiners/taskreports/model"
	"fmt"
	"log"
	"net/http"
	"bitbucket.org/new-joiners/taskreports/controller"
	"bitbucket.org/new-joiners/taskreports/platform"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(config model.ApplicationConfig) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(fmt.Sprintf("%v/", config.ContextPath), homePage)
	router.HandleFunc(fmt.Sprintf("%v/status",  config.ContextPath), controller.GetStatus).Methods("GET")

	log.Printf("Starting service with port %v and context path %v", config.Port, config.ContextPath)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), router))
}

func main() {
	config := platform.GetConfiguration()
	platform.ConfigureContainer(config)
	handleRequests(config)
}