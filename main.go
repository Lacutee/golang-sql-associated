package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RouterUse(router *mux.Router) {
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode("Hello World")
	})
}

func RouterStart() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println(`Running on port 8090`)
	RouterUse(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func main() {
	RouterStart()
}
