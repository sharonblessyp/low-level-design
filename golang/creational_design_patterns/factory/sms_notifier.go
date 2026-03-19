package factory

import "fmt"

type SMSNotification struct{}

func GetSMSNotifier() *SMSNotification {
	return &SMSNotification{}
}

func (sms *SMSNotification) send(message string) {
	fmt.Printf("Sending SMS: %s", message)
}
