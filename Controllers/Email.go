package Controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/smtp"
)

func Email(c *gin.Context) {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "gits5622@gmail.com", "gitau1998", "smtp.mailtrap.io")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"ggitau5622@gmail.com"}
	msg := []byte("To: ggitau5622@gmail.com\r\n" +
		"Subject: Why are you not using Mailtrap yet?\r\n" +
		"\r\n" +
		"Here’s the space for our great sales pitch\r\n")
	err := smtp.SendMail("smtp.mailtrap.io:25", auth, "piotr@mailtrap.io", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
