package factory

import "fmt"

type PushNotification struct{}

func (p *PushNotification) send(message string) {
	fmt.Printf("Sending push notification: %s", message)
}
