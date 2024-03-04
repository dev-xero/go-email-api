package handler

import (
	"dev-xero/email-api/util"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type EmailHandler struct{}

func (handler *EmailHandler) Home(w http.ResponseWriter, r *http.Request) {
	response := util.Response{
		Message: "Email endpoint requested.",
		Code:    http.StatusOK,
		Success: true,
		Payload: nil,
	}

	response.SendEncoded(w)
}

func (handler *EmailHandler) Send(w http.ResponseWriter, r *http.Request) {
	passkey := os.Getenv("EMAIL_PASSKEY")
	log.Println("[INFO]: Pass Key:", passkey)

	// Get the sender and recipient email addresses
	emailRequest := util.EmailRequest{}
	json.NewDecoder(r.Body).Decode(&emailRequest)

	log.Println("[INFO]: Sender:", emailRequest.Sender, "Recipient:", emailRequest.Recipient)

	if emailRequest.Sender == "" || emailRequest.Recipient == "" {
		response := &util.Response{
			Message: "Invalid sender or recipient received.",
			Code:    http.StatusBadRequest,
			Success: false,
			Payload: nil,
		}
		response.SendEncoded(w)
		return
	}

	// Verification email
	email := util.Email{
		Subject: "Your Verification Code",
		Code:    "7074",
	}

	err := email.Send(emailRequest.Sender, emailRequest.Recipient)

	if err != nil {
		response := &util.Response{
			Message: "Unable to send verification email.",
			Code:    http.StatusInternalServerError,
			Success: false,
			Payload: nil,
		}
		response.SendEncoded(w)
		return
	}

	response := &util.Response{
		Message: "Email sent successfully.",
		Code:    http.StatusOK,
		Success: true,
		Payload: nil,
	}

	response.SendEncoded(w)
}
