package main

import (
	"fmt"
	"net/http"

	"github.com/IGGIAUS/Go-WebServer/web"
	"github.com/gorilla/mux"
)

const (
	PORT = "8000"
)

// Main entry point for the application
func main() {

	// Create the request mux
	var mux *mux.Router = mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/", homeHandler)
	mux.NotFoundHandler = new(web.CustomNotFoundHandler)

	// Configure routes
	routeConfig(mux)

	// Start the web server
	fmt.Println("Starting the web server...")
	err := http.ListenAndServe(":"+PORT, mux)
	if err != nil {
		panic("Start failed: " + err.Error())
	}
	fmt.Println("Server Runing on port: " + PORT)
}

// Here all of the routes are registered
func routeConfig(mux *mux.Router) {
	//Register webservice sections
}

// Handler for the root of the webservices "/"
func homeHandler(writer http.ResponseWriter, request *http.Request) {
	var message string = " ____        _ _     _ _               ____  _            _\n|  _ \\      (_) |   | (_)             |  _ \\| |          | |\n| |_) |_   _ _| | __| |_ _ __   __ _  | |_) | | ___   ___| | _____\n|  _ <| | | | | |/ _` | | '_ \\ / _` | |  _ <| |/ _ \\ / __| |/ / __|\n| |_) | |_| | | | (_| | | | | | (_| | | |_) | | (_) | (__|   <\\__ \\\n|____/ \\__,_|_|_|\\__,_|_|_| |_|\\__, | |____/|_|\\___/ \\___|_|\\_\\___/\n                                __/ |\n                               |___/\n                          --[WEB API]--\n"
	writer.Write([]byte(message))
}
