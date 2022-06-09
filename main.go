package main

import (
	"log"
	"os"

	"github.com/talbx/celepush/internal/pkg/delegate"
	"github.com/talbx/celepush/internal/pkg/model"
	"github.com/talbx/celepush/internal/pkg/service"
)

func main() {
	log.Println("celepush initialized!")
	// CheckBirthdays()
	log.Println("celepush terminating!")
	os.Exit(0)
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
