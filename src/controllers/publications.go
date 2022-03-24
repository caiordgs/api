package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreatePost creates a post on the database.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publications models.Publications
	if erro = json.Unmarshal(requestBody, &publications); erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	publications.AuthorID = userID

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewPublicationsRepository(db)
	publications.ID, erro = repository.Create(publications)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, publications)
}

// FindPosts brings up the posts that show up on the user feed.
func FindPosts(w http.ResponseWriter, r *http.Request) {

}

// FindPost brings up a specific post.
func FindPost(w http.ResponseWriter, r *http.Request) {

}

// UpdatePost modifies a post.
func UpdatePost(w http.ResponseWriter, r *http.Request) {

}

// DeletePost deletes a post.
func DeletePost(w http.ResponseWriter, r *http.Request) {

}
