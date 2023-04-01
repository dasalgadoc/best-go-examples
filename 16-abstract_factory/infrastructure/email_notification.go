package infrastructure

import (
	"dasalgadoc.com/best-go-examples/16-abstract_factory/domain"
	"fmt"
)

type EmailNotification struct {
}

func NewEmailNotification() domain.NotificationFactory {
	return &EmailNotification{}
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification by Email...")
}

func (EmailNotification) GetSender() domain.Sender {
	return EmailNotificationSender{}
}
