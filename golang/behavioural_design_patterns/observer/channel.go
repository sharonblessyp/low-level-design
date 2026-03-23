package observer

type Channel interface {
	Subsribe()
	UnSubsribe()
	Notify()
}
