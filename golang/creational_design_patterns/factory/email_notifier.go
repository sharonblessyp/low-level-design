package factory

import "fmt"

type EmailNotification struct{}

func GetEmailNotifier() *EmailNotification {
	return &EmailNotification{}
}

func (e *EmailNotification) send(message string) {
	fmt.Printf("Sending push notification: %s", message)
}
