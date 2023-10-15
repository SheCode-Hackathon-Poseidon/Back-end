package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	// "sample/api/auth"
	"sample/api/exitcode"
	formaterror "sample/api/utils/errors"

	"sample/api/models"

	"sample/api/responses"
)

// CreateUser is...
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user.Prepare()

	err = user.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, formattedError)

		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

	responses.JSON(w, http.StatusCreated, userCreated)
}

// GetUsers is...
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, err)

		return
	}

	if len(*users) == 0 {
		responses.ERROR(w, http.StatusNotFound, exitcode.BE_FAILED, errors.New("No users found"))

		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUser is...
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, err)

		return
	}

	responses.JSON(w, http.StatusOK, userGotten)
}

// UpdateUser is...
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	// tokenID, err := auth.ExtractTokenID(r)

	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

	// 	return
	// }

	// if tokenID != uint32(uid) {
	// 	responses.ERROR(w, http.StatusBadRequest, exitcode.BE_FAILED, errors.New("ID do Token inválido"))
	// }

	user.Prepare()

	err = user.Validate("update")

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, err)

		return
	}

	updatedUser, err := user.UpdateUser(server.DB, uint32(uid))

	if err != nil {
		formatedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, formatedError)

		return
	}

	responses.JSON(w, http.StatusOK, updatedUser)
}

// DeleteUser is...
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user := models.User{}

	// tokenID, err := auth.ExtractTokenID(r)

	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, err)

	// 	return
	// }

	// if tokenID != uint32(uid) {
	// 	responses.ERROR(w, http.StatusForbidden, exitcode.BE_FAILED, errors.New("Token invalid"))

	// 	return
	// }

	_, err = user.DeleteUser(server.DB, uint32(uid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, err)

		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", uid))

	responses.JSON(w, http.StatusNoContent, "")
}
