package notificator

import (
	"fmt"
	"os"

	"github.com/vonage/vonage-go-sdk"
)

func NotifyIfNotSoldOut() {
	API_KEY := os.Getenv("SMS_KEY")
	API_SECRET := os.Getenv("SMS_SECRET")

	number := "50672335061"

	auth := vonage.CreateAuthFromKeySecret(API_KEY, API_SECRET)
	smsClient := vonage.NewSMSClient(auth)
	response, errResp, err := smsClient.Send(number, number, "THE KINESIS KEYBORD IS AVAILABLE", vonage.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Message sent")

		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	} else {
		fmt.Println("Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
	}
}
