package utility

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/p7chkn/go_api/users"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/all", users.ReturnUsers)
	log.Fatal(http.ListenAndServe(":100000", myRouter))
}
