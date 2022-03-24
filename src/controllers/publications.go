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

	var posts models.Posts
	if erro = json.Unmarshal(requestBody, &posts); erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	posts.AuthorID = userID

	if erro = posts.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewPublicationsRepository(db)
	posts.ID, erro = repository.Create(posts)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, posts)
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
