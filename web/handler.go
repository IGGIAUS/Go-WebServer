package web

import (
	"encoding/json"
	"net/http"

	"github.com/IGGIAUS/Webservices/web"
)

// Handles the JSON web service request for a specific http method.
// Returns the status code and the response object to be serialized into JSON.
type RequestHandler func(request *http.Request) (object interface{}, statusCode int)

// Custom handler for when a request is not found
type CustomNotFoundHandler struct{}

// Custom Handler that separates different methods into their own handler functions
type HttpMethodHandler struct {
	Get     RequestHandler
	Put     RequestHandler
	Post    RequestHandler
	Patch   RequestHandler
	Delete  RequestHandler
	Options RequestHandler
}

// The function that is called on the handler when a request arrives that needs to be handled
func (handler *HttpMethodHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	web.LogRequest(request)
	// TODO: AuthKey
	//authkey := GetUrlParameter(request, "access_token")

	//headers
	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", allowedMethords(handler))

	// Check the incomming request for the "application/json" header value
	if request.Header["Content-Type"][0] != "application/json" {
		writer.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	// Find the handler function that'll get called
	var handlerFunc RequestHandler = nil
	switch request.Method {
	case "GET":
		handlerFunc = handler.Get
	case "PUT":
		handlerFunc = handler.Put
	case "POST":
		handlerFunc = handler.Post
	case "PATCH":
		handlerFunc = handler.Patch
	case "DELETE":
		handlerFunc = handler.Delete
	case "OPTIONS":
		handlerFunc = handler.Options
	}

	// If not valid handler was found then return a MethodNotAllowed
	if handlerFunc == nil {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		// Call the handler
		object, statusCode := handlerFunc(request)
		var jsonData []byte
		if Str, ook := object.(string); ook == true {
			jsonData = []byte(Str)
		} else {
			// Serialze the object to JSOn
			if object != nil {
				jsonData = toJson(object)
			}
		}

		// Set the response status code
		writer.WriteHeader(statusCode)

		// Set the body of the request
		writer.Write(jsonData)
	}
}

func (handler *CustomNotFoundHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	web.LogRequest(request)
	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte("{\"Error\": \"404 - not found\"}"))
}

//Returns a string in the requred format for allowed method
func allowedMethords(handler *HttpMethodHandler) string {
	methods := ""
	commer := false
	if handler.Get != nil {
		methods += "GET, HEAD"
		commer = true
	}
	if handler.Post != nil {
		if commer == true {
			methods += ", "
		} else {
			commer = true
		}
		methods += "POST"
	}
	if handler.Put != nil {
		if commer == true {
			methods += ", "
		} else {
			commer = true
		}
		methods += "PUT"
	}
	if handler.Delete != nil {
		if commer == true {
			methods += ", "
		} else {
			commer = true
		}
		methods += "DELETE"
	}
	if handler.Patch != nil {
		if commer == true {
			methods += ", "
		}
		methods += "PATCH"
	}
	if handler.Options != nil {
		if commer == true {
			methods += ", "
		}
		methods += "OPTIONS"
	}
	return methods
}

// Converts a struct to JSON and returns it. If the conversion fails then an
// empty string is returned instead
func toJson(object interface{}) []byte {
	jsonBytes, _ := json.Marshal(object)
	return jsonBytes
}
