package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"sample/api/auth"
	"sample/api/exitcode"
	"sample/api/models"
	"sample/api/responses"
	formaterror "sample/api/utils/errors"
)

// Login is...
func (server *Server) Login(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	user.Prepare()
	err = user.Validate("login")

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(res, http.StatusBadRequest, exitcode.WRONG_ACCOUNT_CREDENTIAL, formattedError)
		return
	}

	_, err = server.SaveToken(user.Email, token)

	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, exitcode.WRONG_ACCOUNT_CREDENTIAL, err)
		return
	}

	responses.JSON(res, http.StatusOK, map[string]interface{}{
		"exitcode": 0,
		"message":  "Login success",
		"token":    token,
	})
}

// SignIn is...
func (server *Server) SignIn(email, password string) (string, error) {
	user := models.User{}

	err := server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", errors.New("email or password does not match")
	}

	match := models.VerifyPassword(user.Password, password)

	if !match {
		return "", errors.New("email or password does not match")
	}

	return auth.CreateToken(user.ID)
}

// Save token into db
func (server *Server) SaveToken(email, token string) (string, error) {
	err := server.DB.Debug().Model(models.User{}).Where("email = ?", email).Update(map[string]interface{}{
		"token": token,
	}).Error

	return "", err
}
