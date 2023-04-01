package domain

type Sender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}
