package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/register", RegisterEndPoint).Methods("POST")
	router.HandleFunc("/gen-token", GenTokenEndPoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func RegisterEndPoint(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Not implemented yet!")
	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	var user User

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error in request.\nWe need json with \"username\" and \"password\" fields.\n")
		return
	}

	RegisterUser(w, user.Username, user.Password)
}

func GenTokenEndPoint(w http.ResponseWriter, r *http.Request) {

}

func GenJokeEndPoint(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://geek-jokes.sameerkumar.website/api")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)
	//SaveJoke(token, string(body))

}

func SaveJoke(token, joke string) {
	fmt.Println("Not implemented yet!")
}

func AllJokesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet!")
}

func DeleteJokeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet!")
}

func UpdateJokeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not implemented yet!")
}
