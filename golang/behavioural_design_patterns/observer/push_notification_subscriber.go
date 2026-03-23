package observer

import "fmt"

type PushNotification struct {
	deviceId string
}

func NewPushNotification(deviceId string) *PushNotification {
	return &PushNotification{
		deviceId: deviceId,
	}
}

func (p *PushNotification) PushNotification() {
	fmt.Printf("push notification sent to %s", p.deviceId)
}

func (p *PushNotification) GetId() string {
	return p.deviceId
}
