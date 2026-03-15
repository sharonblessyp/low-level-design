package factory

import "errors"

type NotifierFactory struct{}

func NewNotifierFactory() *NotifierFactory {
	return &NotifierFactory{}
}

// notification service that creates instances
func (nf NotifierFactory) GetNotifier(notificationType string) (Notifier, error) {

	switch notificationType {
	case "SMS":
		return GetSMSNotifier(), nil

	case "EMAIL":
		return GetEmailNotifier(), nil
	default:
		return nil, errors.New("unknown notifier")
	}
}

func GetSMSNotifier() *SMSNotification {
	return &SMSNotification{}
}

func GetEmailNotifier() *EmailNotification {
	return &EmailNotification{}
}
