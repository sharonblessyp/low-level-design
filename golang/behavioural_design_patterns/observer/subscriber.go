package observer

type Subscriber interface {
	UpdateSubscriber()
	GetId() string
}
