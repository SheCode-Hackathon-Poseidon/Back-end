package controllers

import (
	"fmt"
	"net/http"

	"sample/api/responses"
)

// Home is..
func (server *Server) Home(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Response welcome")
	responses.JSON(res, http.StatusOK, "Welcome to our API!")

}
