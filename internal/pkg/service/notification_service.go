package service

import (
	"fmt"
	"github.com/gregdel/pushover"
	"log"
)

type NotificationService interface {
	Notify(msg *pushover.Message) *pushover.Response
}

type PushNotificationService struct{}

func (notificationService PushNotificationService) Notify(msg *pushover.Message) *pushover.Response {

	app := pushover.New(AppConf.Pushover.Apitoken)
	recipient := pushover.NewRecipient(AppConf.Pushover.Usertoken)
	response, err := app.SendMessage(msg, recipient)
	if err != nil {
		log.Panic(err)
	}

	log.Printf(fmt.Sprintf("Sent out pushover message with id %s", response.ID))
	return response
}
