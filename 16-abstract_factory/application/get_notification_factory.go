package application

import (
	"dasalgadoc.com/best-go-examples/16-abstract_factory/domain"
	"dasalgadoc.com/best-go-examples/16-abstract_factory/infrastructure"
	"fmt"
)

var concreteFactories = map[string]func() domain.NotificationFactory{
	"SMS":   infrastructure.NewSMSNotification,
	"EMAIL": infrastructure.NewEmailNotification,
}

func GetNotificationFactory(notificationType string) (domain.NotificationFactory, error) {
	factory, ok := concreteFactories[notificationType]
	if !ok {
		return nil, fmt.Errorf("no notification type")
	}

	return factory(), nil
}
