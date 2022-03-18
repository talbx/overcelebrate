package service

import (
	"github.com/gregdel/pushover"
	"github.com/talbx/birthday-notice/internal/pkg/utils"
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

	utils.Sugar.Infow("Sent out pushover message with",
		"id", response.ID,
	)

	return response
}
