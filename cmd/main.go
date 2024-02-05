package main

import (
	"fmt"
	"net/http"

	"github.com/Abhishek-Jain-1925/Website-Checker/service"
	"github.com/gorilla/mux"
)

var Map1 map[string]string

func main() {

	fmt.Println("Starting Server...")
	r := mux.NewRouter()

	r.HandleFunc("/websites", service.WebsiteHandler).Methods(http.MethodPost)
	r.HandleFunc("/websites", service.WebsiteStatuses).Methods(http.MethodGet)
	r.HandleFunc("/website", service.ShowStatus).Methods(http.MethodGet)

	http.ListenAndServe("localhost:3000", r)
	fmt.Println("Server started !!")
}
