package main

import (
	"encoding/json"
	"fmt"
	"golang-sql-associated/database"
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

func InitDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "go_lang_asossicated",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate()
}

func main() {
	InitDB()
	RouterStart()
}
