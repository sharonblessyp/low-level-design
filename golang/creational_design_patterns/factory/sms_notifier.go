package factory

import "fmt"

type SMSNotification struct{}

func (sms *SMSNotification) send(message string) {
	fmt.Printf("Sending SMS: %s", message)
}
