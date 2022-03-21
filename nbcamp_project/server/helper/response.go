package helper

import (
	"encoding/json"
	"errors"
	"log"
	apps "nbcamp_project/server/apps/api"
	"net/http"
)

// Handle method not allowed
func HandleNotMethodAllowed(w http.ResponseWriter, method string) {
	handleResponseError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
}

// handle bad request
func HandleBadRequest(w http.ResponseWriter, err error) {
	handleResponseError(w, err, http.StatusBadRequest)
}

// handle internal server error
func HandleInternalServerError(w http.ResponseWriter, err error) {
	handleResponseError(w, err, http.StatusInternalServerError)
}

// handle not found
func HandleNotFound(w http.ResponseWriter, err error) {
	handleResponseError(w, err, http.StatusNotFound)
}

// function for wrapping all handle error response
func handleResponseError(w http.ResponseWriter, err error, status int) {
	log.Println(err)
	// using apps response for API
	var response apps.ResponseFail
	response.Status = status
	response.Message = err.Error()

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}

// handle success, with status code 200
func HandleSuccess(w http.ResponseWriter, payload interface{}) {
	handleResponseSuccess(w, payload, http.StatusOK)
}

// handle created success, with status code 201
func HandleCreateSuccess(w http.ResponseWriter, payload interface{}) {
	handleResponseSuccess(w, payload, http.StatusCreated)
}

// function for wrapping all success response
func handleResponseSuccess(w http.ResponseWriter, payload interface{}, status int) {
	// using apps response for API
	var response apps.ResponseSuccess
	response.Status = status
	response.Data = payload

	w.WriteHeader(response.Status)
	json.NewEncoder(w).Encode(response)
}
