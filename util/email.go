package util

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"gopkg.in/gomail.v2"
)

type Email struct {
	Subject string
	Code    string
}

func (email *Email) readTemplate(path string) (string, error) {
	template, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	if err = template.Execute(buffer, email); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (email *Email) absoluteFilePath(path string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	absPath := filepath.Join(cwd, "html", path)
	return absPath, nil
}

func (email *Email) Send(from string, to string) error {
	path, err := email.absoluteFilePath("template.html")
	if err != nil {
		return err
	}

	template, err := email.readTemplate(path)
	log.Println("[INFO]: Absolute Path:", path)

	if err != nil {
		log.Println("[ERROR]: Unable to read template file:", err)
		return err
	}

	log.Println("[INFO]: Template:", template)

	mail := gomail.NewMessage()

	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", email.Subject)
	mail.SetBody("text/html", template)

	// SMTP credentials
	var (
		host    = os.Getenv("EMAIL_HOST")
		port, _ = strconv.Atoi(os.Getenv("EMAIL_PORT"))
		address = os.Getenv("EMAIL_ADDRESS")
		passkey = os.Getenv("EMAIL_PASSKEY")
	)

	dialer := gomail.NewDialer(host, port, address, passkey)
	dialer.DialAndSend(mail)

	return nil
}
