package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repos.NewUserRepository(db)
	repository.Create(user)
}

// Searches all users.
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users!"))
}

// Searches one specific user.
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching an user!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user!"))
}
