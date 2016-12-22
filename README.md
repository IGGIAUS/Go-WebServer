#Go-WebServer Base
Very Basic Webserver for building JSON web-services on top of.
Created for personal use as I often use Web-services in projects and it would be useful to base I could reuse.

Defaults to port 8000 can be changed by changing the constant in main.go

To add you functions to the server you call Register on the mux in the routeConfig method in main.go.

E.G:
mux.Register(path String, handler web.HttpMethodHandler)

For more information go to https://github.com/gorilla/mux

You can add your functions to the handler providing they match the type of web.RequestHandler.
The handler can have these functions assigned to any of its below properties
