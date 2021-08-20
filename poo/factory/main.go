package main

import "fmt"

type INotifyFactory interface {
	SendNotify()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotify struct {
}

func (SMSNotify) SendNotify() {
	fmt.Println("Sending notification via SMS")
}

type SMSNotifySender struct {
}

func (SMSNotify) GetSender() ISender {
	return SMSNotifySender{}
}

func (SMSNotifySender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotifySender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotify struct {
}

func (EmailNotify) SendNotify() {
	fmt.Println("Send notification via Email")
}

type EmailNotifySender struct {
}

func (EmailNotify) GetSender() ISender {
	return EmailNotifySender{}
}

func (EmailNotifySender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotifySender) GetSenderChannel() string {
	return "Email"
}

func getNotifyFactory(sender string) (INotifyFactory, error) {
	switch sender {
	case "SMS":
		return SMSNotify{}, nil
	case "Email":
		return EmailNotify{}, nil
	}
	return nil, fmt.Errorf("NotifyFactory not found")
}

func sendNotification(f INotifyFactory) {
	f.SendNotify()
}

func getMethod(f INotifyFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	factory, err := getNotifyFactory("SMS")
	if err != nil {
		fmt.Println(err)
	}
	sendNotification(factory)
	getMethod(factory)

	Emailfactory, Emailerr := getNotifyFactory("Email")
	if Emailerr != nil {
		fmt.Println(Emailerr)
	}
	sendNotification(Emailfactory)
	getMethod(Emailfactory)
}
