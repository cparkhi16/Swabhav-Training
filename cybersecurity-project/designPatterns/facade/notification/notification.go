package notification

import "fmt"

type Notification struct {
	msg string
}

func New(msg string) *Notification {
	return &Notification{
		msg: msg,
	}
}

func (n *Notification) SendNotification(msg string) {
	n.msg = msg
	fmt.Println("Notification:-", msg)
}
