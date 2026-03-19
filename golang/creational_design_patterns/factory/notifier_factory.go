package factory

import "errors"

// factory function that returns the object created
func GetNotifier(notificationType string) (Notifier, error) {
	switch notificationType {
	case "SMS":
		return GetSMSNotifier(), nil
	case "EMAIL":
		return GetEmailNotifier(), nil
	case "PUSH":
		return GetPushNotifier(), nil
	default:
		return nil, errors.New("unknown notifier")
	}
}

type NotifierFactory struct{}

func NewNotifierFactory() *NotifierFactory {
	return &NotifierFactory{}
}

func (nf *NotifierFactory) Notify(channel string, message string) error {
	// instance creation
	notifier, err := GetNotifier(channel)
	if err != nil {
		return err
	}
	notifier.send(message)
	return nil
}
