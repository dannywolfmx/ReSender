package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/dannywolfmx/ReSender/app/domain/model"
	"github.com/dannywolfmx/ReSender/app/registry"
	"github.com/gorilla/mux"
)

func sendMail(router *mux.Router, ctn *registry.Container) {
	router.HandleFunc("/send", func(w http.ResponseWriter, h *http.Request) {
		//Get mailserver account
		w.Write([]byte("Message sende"))
	})
}

func send(server model.MailServer, subject, body string, to []string) {
	msg := getMsg(server.From, strings.Join(to, ","), subject, body)
	//Ejemplo smtp.gmail.com:587
	auth := smtp.PlainAuth("", server.From, server.Pass, server.Server)
	err := smtp.SendMail(server.Address, auth, server.From, to, []byte(msg))
	if err != nil {
		log.Println("Smtp error: ", err)
		return
	}
	log.Print("Enviado ")
}

func getMsg(from, to, subject, body string) string {
	return fmt.Sprintf("From: %s\n To: %s\n, Subject: %s\n\n%s", from, to, subject, body)
}
