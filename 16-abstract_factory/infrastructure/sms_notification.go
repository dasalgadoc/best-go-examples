package infrastructure

import (
	"dasalgadoc.com/best-go-examples/16-abstract_factory/domain"
	"fmt"
)

type SMSNotification struct{}

func NewSMSNotification() domain.NotificationFactory {
	return &SMSNotification{}
}

func (s SMSNotification) SendNotification() {
	fmt.Println("Sending notification by SMS...")
}

func (s SMSNotification) GetSender() domain.Sender {
	return SMSNotificationSender{}
}
