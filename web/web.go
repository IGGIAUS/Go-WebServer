package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Standardised error response
type ErrorResponse struct {
	Error string
}

// Logs a request to the standard output stream
func LogRequest(request *http.Request) {
	// Print the time, method and un-parsed URL
	time := time.Now().UTC().Format(time.RFC3339)
	fmt.Printf("%s [%s] %s\n", time, request.Method, request.RequestURI)
}

// Returns all of the bytes in the body of a Request
// If this fails then nil will be returned instead (Empty slice)
func GetBody(request *http.Request) []byte {
	// Get all of the bytes in the request
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil
	} else {
		return bytes
	}
}

// Returns an Ok status code with a JSON payload response (Code: 200)
func Ok(object interface{}) (interface{}, int) {
	return object, http.StatusOK
}

// Returns a NoContent status code with an error message (Code: 204)
func NoContent(errorMessage string) (interface{}, int) {
	return ErrorResponse{Error: errorMessage}, http.StatusNoContent
}

// Returns a BadRequest with an error message (Code: 400)
func BadRequest(errorMessage string) (interface{}, int) {
	return ErrorResponse{Error: errorMessage}, http.StatusBadRequest
}

// Returns an InternalServerError status code with an error message (Code: 500)
func InternalServerError(errorMessage string) (interface{}, int) {
	return ErrorResponse{Error: errorMessage}, http.StatusInternalServerError
}

// Returns a NotImplemented status code with an error message (Code 501)
func NotImplemented(errorMessage string) (interface{}, int) {
	return ErrorResponse{Error: errorMessage}, http.StatusNotImplemented
}
