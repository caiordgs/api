package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("register"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

// SearchUsers searches all users.
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	users, erro := repository.Search(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// SearchUser searches one specific user.
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	user, erro := repository.IdSearch(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser updates the info on a user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != tokenUserID {
		responses.Erro(w, http.StatusForbidden, errors.New("it's not possible to edit an user that's not yours"))
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("edicao"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repos := repos.NewUserRepository(db)
	if erro = repos.Update(userID, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser deletes the user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	tokenUserID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != tokenUserID {
		responses.Erro(w, http.StatusForbidden, errors.New("you're not allowed to delete an user that's not yours"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	if erro = repository.Delete(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// FollowUser follows another user.
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	if erro = repository.Follow(userID, followerID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser unfollows another user.
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("you can't unfollow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	if erro = repository.Unfollow(userID, followerID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FindFollowers brings up all followers from a user.
func FindFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	followers, erro := repository.FindFollowers(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// FindFollowing brings up all other users that a user is following.
func FindFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	users, erro := repository.FindFollowing(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// PasswordUpdate allows to change the password of a user.
func PasswordUpdate(w http.ResponseWriter, r *http.Request) {
	tokenUserID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if tokenUserID != userID {
		responses.Erro(w, http.StatusForbidden, errors.New("it's not possible to update the password from other users"))
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)

	var password models.Password
	if erro = json.Unmarshal(requestBody, &password); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repos.NewUserRepository(db)
	savedPassword, erro := repository.SearchPassword(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.PasswordValidation(savedPassword, password.Old); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("your old password does not match"))
		return
	}

	hashedPassword, erro := security.Hash(password.New)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.PasswordUpdate(userID, string(hashedPassword)); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
