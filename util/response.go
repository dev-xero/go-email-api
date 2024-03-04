package util

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

func (res *Response) SendEncoded(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)

	json.NewEncoder(w).Encode(res)
}
