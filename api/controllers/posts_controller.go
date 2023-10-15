package controllers

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"sample/api/exitcode"
	"sample/api/models"
	"sample/api/responses"
	formaterror "sample/api/utils/errors"
)

// CreatePost is...
func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)
		return
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)
		return
	}

	project := models.Project{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		ShareMode:   data["share_mode"].(string),
		Status:      data["status"].(string),
	}

	token, ok := data["token"].(string)
	if !ok {
		responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)
		return
	}

	// tasks, ok := data["tasks"].([]interface{})
	// if !ok {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)
	// 	return
	// }

	project.Prepare()
	// fmt.Printf("Create post | data =  %+v; token =  %s; tasks = %+v", data, token, tasks)

	user := models.User{}

	userGotten, err := user.FindUserByToken(server.DB, token)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, exitcode.BE_FAILED, errors.New("invalid token"))
		return
	}

	project.UserId = int(userGotten.ID)

	projectCreated, err := project.Create(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, exitcode.BE_FAILED, err)
		return
	}

	// for _, taskData := range tasks {
	// 	task, ok := taskData.(map[string]interface{})

	// 	if !ok {
	// 		fmt.Println("Invalid task data")
	// 		continue
	// 	}

	// 	newTask := models.Task{
	// 		Title:     task["title"].(string),
	// 		Deadline:  task["deadline"].(string),
	// 		ProjectID: int(projectCreated.ID),
	// 	}

	// 	// Create the task in the database
	// 	newTask.Create(server.DB)

	// 	// Print added task details (optional)
	// 	fmt.Printf("Added Task:  %+v", newTask)
	// }

	responses.JSON(w, http.StatusCreated, projectCreated)
}

// GetPosts is...
func (server *Server) GetPosts(res http.ResponseWriter, req *http.Request) {
	var projects []models.Project

	if err := server.DB.Preload("Tasks").Find(&projects).Error; err != nil {
		responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, err)
		return
	}

	if len(projects) == 0 {
		responses.ERROR(res, http.StatusNotFound, exitcode.BE_FAILED, errors.New("No projects found"))
		return
	}

	responses.JSON(res, http.StatusOK, projects)
}

// GetPost is...
func (server *Server) GetPost(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, err)

		return
	}

	post := models.Project{}

	foundPost, err := post.FindPostByID(server.DB, uint32(pid))

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	responses.JSON(res, http.StatusOK, foundPost)
}

// UpdatePost is...
func (server *Server) UpdatePost(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	// tokenID, err := auth.ExtractTokenID(req)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	post := models.Project{}
	foundPost, err := post.FindPostByID(server.DB, uint32(pid))

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	// if tokenID != foundPost.AuthorID {
	// 	responses.ERROR(res, http.StatusUnauthorized, exitcode.BE_FAILED, errors.New("Only the author can update the post"))

	// 	return
	// }

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	postUpdate := models.Project{}

	err = json.Unmarshal(body, &postUpdate)

	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

		return
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()

	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, exitcode.BE_FAILED, err)

		return
	}

	postUpdate.ID = foundPost.ID

	postUpdated, err := postUpdate.UpdatePost(server.DB)

	if err != nil {
		formatedError := formaterror.FormatError(err.Error())

		responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, formatedError)

		return
	}

	responses.JSON(res, http.StatusOK, postUpdated)
}

// DeletePost is...
func (server *Server) DeletePost(res http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(req)

	// pid, err := strconv.ParseUint(vars["id"], 10, 64)

	// if err != nil {
	// 	responses.ERROR(res, http.StatusBadRequest, exitcode.BE_FAILED, err)

	// 	return
	// }

	// post := models.Project{}
	// foundPost, err := post.FindPostByID(server.DB, uint32(pid))

	// if err != nil {
	// 	responses.ERROR(res, http.StatusUnprocessableEntity, exitcode.BE_FAILED, err)

	// 	return
	// }

	// tokenID, err := auth.ExtractTokenID(req)

	// if err != nil {
	// 	responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, err)

	// 	return
	// }

	// if tokenID != foundPost.AuthorID {
	// 	responses.ERROR(res, http.StatusUnauthorized, exitcode.BE_FAILED, errors.New("Only the author can delete the post"))

	// 	return
	// }

	// _, err = foundPost.DeletePost(server.DB, foundPost.ID, tokenID)

	// if err != nil {
	// 	formatedError := formaterror.FormatError(err.Error())

	// 	responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, formatedError)

	// 	return
	// }

	responses.JSON(res, http.StatusNoContent, "")
}
