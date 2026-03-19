package factory

import "fmt"

type PushNotification struct{}

func GetPushNotifier() *PushNotification {
	return &PushNotification{}
}

func (p *PushNotification) send(message string) {
	fmt.Printf("Sending push notification: %s", message)
}
