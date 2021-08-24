package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Response represent response being sent
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"Data"`
}

//Err represents erroe being sent
type Err struct {
	Error        error       `json:"error"`
	Message      string      `jason:"message"`
	Data         interface{} `jason:"data"`
	ResponseCode int         `json:"response_code"`
}

//WriteSuccess which returns successful response
func WriteSuccess(w http.ResponseWriter, responseCode int, message string, data interface{}) {
	var r Response
	r.Success = true
	r.Data = data
	r.Message = message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(r)

}

//WriteError returns error response
func WriteError(w http.ResponseWriter, err Err) {
	fmt.Println(err.Message, err.Error)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.ResponseCode)
	json.NewEncoder(w).Encode(err)

}
