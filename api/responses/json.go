package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON is...
func JSON(res http.ResponseWriter, statusCode int, data interface{}) {
	res.WriteHeader(statusCode)

	err := json.NewEncoder(res).Encode(data)

	if err != nil {
		fmt.Fprintf(res, "%s", err.Error())
	}
}

// ERROR is..
func ERROR(res http.ResponseWriter, statusCode int, exitcode int, err error) {
	if err != nil {
		JSON(res, statusCode, struct {
			Error string `json:"error"`
			Exitcode int  `json:"exitcode"`
		}{
			Error: err.Error(),
			Exitcode: exitcode,
		})

		return
	}

	fmt.Printf("Error: %s; Exitcode = %d", err.Error(), exitcode)

	JSON(res, http.StatusBadRequest, nil)
}
