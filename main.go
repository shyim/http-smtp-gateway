package main

import (
	"encoding/json"
	"github.com/xhit/go-simple-mail/v2"
	"io"
	"log"
	"net/http"
	"os"
)

type MailMessage struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func main() {
	authToken := os.Getenv("AUTH_TOKEN")

	server := mail.NewSMTPClient()

	server.Host = os.Getenv("SMTP_HOST")
	server.Port = 587
	server.Username = os.Getenv("SMTP_USER")
	server.Password = os.Getenv("SMTP_PASS")
	server.Encryption = mail.EncryptionTLS

	server.Authentication = mail.AuthPlain

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != authToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		content, err := io.ReadAll(r.Body)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var message MailMessage
		err = json.Unmarshal(content, &message)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		smtpClient, err := server.Connect()

		if err != nil {
			log.Fatal(err)
		}

		email := mail.NewMSG()
		email.SetFrom(message.From).
			AddTo(message.To).
			SetSubject(message.Subject)

		email.SetBody(mail.TextHTML, message.Body)

		go func() {
			err = email.Send(smtpClient)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Email Sent")
			}
		}()
		w.WriteHeader(http.StatusNoContent)
	})

	log.Printf("Listing for requests at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
