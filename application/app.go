package application

import (
	"dev-xero/email-api/route"
	"dev-xero/email-api/util"
	"net/http"
)

type Application struct {
	mux *http.ServeMux
}

// Creates a new Application instance
func (app *Application) New(multiplexer *http.ServeMux) {
	app.mux = multiplexer
}

// Initializes the application and listens for requests
func (app *Application) Initialize() {
	route.GetApplicationRoutes(app.mux)

	// Listen for requests on the root endpoint
	app.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			response := util.Response{
				Message: "Invalid endpoint requested.",
				Code:    http.StatusNotFound,
				Success: false,
				Payload: nil,
			}
			response.SendEncoded(w)
			return
		}
		response := util.Response{
			Message: "Welcome to the email API.",
			Code:    http.StatusOK,
			Success: true,
			Payload: nil,
		}
		response.SendEncoded(w)
	})
}
