package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! I'm future Simple API")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandle)

	log.Fatal(http.ListenAndServe(":8081", router))
}
