package infrastructure

type SMSNotificationSender struct{}

func (s SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (s SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
