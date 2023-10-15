package controllers

import (
	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"sample/api/exitcode"
	"sample/api/models"
	"sample/api/responses"
)

func (server *Server) GetProgress(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	pid, err := strconv.ParseUint(vars["project_id"], 10, 32)

	if err != nil {
		responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, err)
		return
	}

	var progress models.Progress

	foundProgress, err := progress.FindProgressOfProject(server.DB, uint32(pid))

    if  err != nil {
        responses.ERROR(res, http.StatusInternalServerError, exitcode.BE_FAILED, err)
        return
    }

    responses.JSON(res, http.StatusOK, foundProgress)
}
