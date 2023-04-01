package main

import (
	"dasalgadoc.com/best-go-examples/16-abstract_factory/application"
	"dasalgadoc.com/best-go-examples/16-abstract_factory/domain"
)

func main() {
	smsFactory, _ := application.GetNotificationFactory("SMS")
	emailFactory, _ := application.GetNotificationFactory("EMAIL")

	domain.Send(smsFactory)
	domain.GetMethod(smsFactory)

	domain.Send(emailFactory)
	domain.GetMethod(emailFactory)

}
