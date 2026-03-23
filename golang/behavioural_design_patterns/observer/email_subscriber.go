package observer

import "fmt"

type EmailSubscriber struct {
	email string
}

func NewEmailSubscriber(email string) *EmailSubscriber {
	return &EmailSubscriber{
		email: email,
	}
}

func (e *EmailSubscriber) UpdateSubscriber() {
	fmt.Printf("email sent to %s \n", e.email)
}

func (e *EmailSubscriber) GetId() string {
	return e.email
}
