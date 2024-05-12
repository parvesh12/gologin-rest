package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/parvesh12/gologin-rest/api"
)

func main() {

	mux := mux.NewRouter()

	//routes
	mux.HandleFunc("/", api.Login).Methods("POST")

	fmt.Printf("Starting server on port:%s\n", os.Getenv("PORT"))

	src := http.Server{
		Addr:        os.Getenv("DOMAIN") + ":" + os.Getenv("PORT"),
		ReadTimeout: time.Second * 10,
		Handler:     mux,
	}

	if err := src.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
