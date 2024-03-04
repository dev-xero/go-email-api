package route

import "net/http"

func GetApplicationRoutes(multiplexer *http.ServeMux) {
	emailMux := GetEmailRoutes()
	multiplexer.Handle("/email/", http.StripPrefix("/email", emailMux))
}
