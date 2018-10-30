package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
