package utility

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnUsers)
	log.Fatal(http.ListenAndServe(":100000", myRouter))
}
