package util

type EmailRequest struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}
