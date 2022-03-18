package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/talbx/birthday-notice/internal/pkg/delegate"
	"github.com/talbx/birthday-notice/internal/pkg/model"
	"github.com/talbx/birthday-notice/internal/pkg/service"
	"github.com/talbx/birthday-notice/internal/pkg/utils"
)

func main() {
	utils.Sugar.Info("birthday-notice initialized!")
	//	_ = gocron.Every(1).Day().At("09:00").Do(CheckBirthdays)
	_ = gocron.Every(1).Minute().Do(CheckBirthdays)
	<- gocron.Start()
}

func CheckBirthdays() {
	utils.Sugar.Infow("Executing birthday check")
	var entries []model.BirthdayEntry
	service.ReadCandidates(&entries)
	processingDelegate := createDelegate()
	processingDelegate.Delegate(&entries)
	utils.Sugar.Infow("Finished execution of birthday check")

}

func createDelegate() delegate.ProcessingDelegate {
	notificationService := service.PushNotificationService{}
	messageDelegate := delegate.MessageDelegateImpl{NotificationService: notificationService}
	processingDelegate := delegate.ProcessingDelegate{MessageDelegate: messageDelegate}
	return processingDelegate
}


