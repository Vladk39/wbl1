package main

import (
	"fmt"
)

type Notifier interface {
	Send(msg msg) error
}

type msg struct {
	message string
	to      string
}

type EmailAdapter struct {
	service *EmailService
}

func (ea *EmailAdapter) Send(msg msg) error {
	return ea.service.SendEmail(msg.to, msg.message)
}

func NotifyUser(n Notifier, msg msg) {
	err := n.Send(msg)
	if err != nil {
		panic("")
	}
}

type EmailService struct{}

func (e *EmailService) SendEmail(to, body string) error {
	fmt.Printf("send mail %s, to %s", body, to)
	return nil
}

func main() {
	EmailService := &EmailService{}

	adapter := &EmailAdapter{
		service: EmailService,
	}

	message := &msg{
		message: "qq",
		to:      "me",
	}

	NotifyUser(adapter, *message)
}
