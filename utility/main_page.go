package utility

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	result := map[string]string{}
	result["result"] = "Welcom to main page!"
	fmt.Println("Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(result)
}
