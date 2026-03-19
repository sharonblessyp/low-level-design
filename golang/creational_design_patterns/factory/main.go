package factory

import "fmt"

/*
1. product interface
2. methods implementig product types
3. factory function -> returns interface
4. call the method based on type sent by client
*/

func main() {
	// create factory
	factory := NewNotifierFactory()

	notificationType := "SMS"
	err := factory.Notify(notificationType, "msg sent")
	if err != nil {
		fmt.Printf("error getting notifier %v", err)
	}
}
