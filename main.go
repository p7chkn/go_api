package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/p7chkn/go_api/users"
	"github.com/p7chkn/go_api/utility"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", utility.HomePage)
	myRouter.HandleFunc("/all", users.ReturnUsers)
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
