package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mail/src/gmail"
	"net/http"

	"github.com/gorilla/mux"
)

// Receiver unexported
type Receiver struct {
	Message string `json:"Message"`
	Email   string `json:"Email"`
	Subject string `json:"Subject"`
}

// SendMail unexported
func SendMail(w http.ResponseWriter, r *http.Request, email string, password string) {
	sender := gmail.Init(email, password)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var receiver Receiver
	json.Unmarshal(reqBody, &receiver)

	fmt.Println(receiver)
	fmt.Println(sender)

	recipients := []string{receiver.Email}
	subject := receiver.Subject
	message := receiver.Message

	bodyMessage := sender.WriteHTMLEmail(recipients, subject, message)

	sender.SendMail(recipients, subject, bodyMessage)
	json.NewEncoder(w).Encode(receiver)
}

// HomePage unexported
func HomePage(w http.ResponseWriter, r *http.Request, appVersion string) {
	fmt.Fprintf(w, "Go Mail Server "+appVersion)
}

// InitRoutes exported
func InitRoutes(port string, email string, password string, appVersion string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomePage(w, r, appVersion)
	}).Methods("GET")

	router.HandleFunc("/api/send/mail", func(w http.ResponseWriter, r *http.Request) {
		SendMail(w, r, email, password)
	}).Methods("POST")

	fmt.Println("App listening on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
