package route

import (
	"dev-xero/email-api/handler"
	"dev-xero/email-api/util"
	"net/http"
)

func GetEmailRoutes() *http.ServeMux {
	emailMux := http.NewServeMux()
	emailHandler := &handler.EmailHandler{}

	emailMux.HandleFunc("POST /send", emailHandler.Send)

	emailMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		emailHandler.Home(w, r)
	})
	return emailMux
}
