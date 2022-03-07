package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))
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
