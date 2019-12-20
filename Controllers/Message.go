package Controllers



import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nexmo-community/nexmo-go"
	"log"
	"net/http"
)

func SendMessages(c *gin.Context) {

	// Auth
	auth := nexmo.NewAuthSet()
	auth.SetAPISecret("5a8f4581", "knlI7i5383mfDWNM")

	// Init Nexmo
	client := nexmo.NewClient(http.DefaultClient, auth)

	// SMS
	smsContent := nexmo.SendSMSRequest{
		From: "Nexmo",
		To:   "254741790736",
		Text: "This is a message sent from Go!.Golang is cool and awesome.Cool kids should consider this!!!",
	}

	smsResponse, _, err := client.SMS.SendSMS(smsContent)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status:", smsResponse.Messages[0].Status)
}
