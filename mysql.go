package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"regexp"
)

// Constants for MYSQL connection
const (
	AddressMySQL  = "127.0.0.1"
	PortMySQL     = "3306"
	UsernameMySQL = "jokeAPI"
	PasswordMySQL = "c7wVzduaTc"
)

/*
Registers user
username - username of account that needs to be registered
password - pass of account that needs to be registered
*/
func RegisterUser(w http.ResponseWriter, username, password string) {
	db, err := sql.Open("mysql", UsernameMySQL+":"+PasswordMySQL+"@tcp("+AddressMySQL+":"+
		PortMySQL+")/JokeAPIDB")
	if err != nil {
		log.Printf("Error when connecting to MySQL\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Creating table for users if it not exists
	query := "CREATE TABLE IF NOT EXISTS users (" +
		"id INT NOT NULL AUTO_INCREMENT, " +
		"username CHAR(18) NOT NULL, " +
		"password VARCHAR(255) NOT NULL, " +
		"PRIMARY KEY (id)" +
		")"

	_, err = db.Query(query)
	if err != nil {
		log.Printf("Error when creating Table users\n%v\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Username validation
	if username != "" && password != "" {
		matched, err := regexp.MatchString("^[a-z0-9_-]{3,16}$", username)
		if err != nil {
			log.Printf("Error when matching regular expression\n%v\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !matched {
			fmt.Fprintln(w, "You need to use username with lowercase letters and length between 3 and 16")
		} else {
			// Checking username existence
			var count int
			row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username)
			err = row.Scan(&count)
			if err != nil {
				log.Printf("Error when counting existing username\n%v\n", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Print(count)
			if count > 0 {
				fmt.Fprintln(w, "This username is already registered. Please, choose another one.")
				return
			}

			// Password validation
			matched, err := regexp.MatchString("^[a-zA-Z0-9_-]{3,16}$", password)
			if err != nil {
				log.Printf("Error when matching regular expression\n%v\n", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !matched {
				fmt.Fprintln(w, "Password length 3-16 is needed\nYou can use the following symbols: [a-z A-Z 0-9 -]")
			} else {
				// Hashing password with md5 and save credentials to database
				hasher := md5.New()
				io.WriteString(hasher, password)
				_, err := db.Query("INSERT INTO users (username, password) VALUES (?, ?)", username,
					hex.EncodeToString(hasher.Sum(nil)))
				if err != nil {
					log.Printf("Error when querying to insert user into DB\n%v\n", err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "Thanks for registration %s\n", username)
			}
		}
	} else {
		fmt.Fprintln(w, "Error in request.\nAPI needs json with \"username\" and \"password\" fields.\n"+
			"You need to use username with lowercase letters and length between 3 and 16.\n"+
			"Password length 3-16 is needed\n"+
			"You can use the following symbols: [a-z A-Z 0-9 -]")
	}
}
