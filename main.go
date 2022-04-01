package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/talbx/celepush/internal/pkg/delegate"
	"github.com/talbx/celepush/internal/pkg/model"
	"github.com/talbx/celepush/internal/pkg/service"
	"log"
)

func main() {
	log.Println("celepush initialized!")
	_ = gocron.Every(1).Day().At("09:00").Do(CheckBirthdays)
	<-gocron.Start()
}

func CheckBirthdays() {
	log.Println("Executing birthday check")
	var entries []model.BirthdayEntry
	service.ReadCandidates(&entries)
	processingDelegate := createDelegate()
	processingDelegate.Delegate(&entries)
	log.Println("Finished execution of birthday check")

}

func createDelegate() delegate.ProcessingDelegate {
	notificationService := service.PushNotificationService{}
	messageDelegate := delegate.MessageDelegateImpl{NotificationService: notificationService}
	processingDelegate := delegate.ProcessingDelegate{MessageDelegate: messageDelegate}
	return processingDelegate
}
