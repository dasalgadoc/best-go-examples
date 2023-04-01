package domain

import "fmt"

type NotificationFactory interface {
	SendNotification()
	GetSender() Sender
}

func Send(n NotificationFactory) {
	n.SendNotification()
}

func GetMethod(n NotificationFactory) {
	fmt.Println(n.GetSender().GetSenderMethod())
}
