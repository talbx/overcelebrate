package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/talbx/celepush/internal/pkg/delegate"
	"github.com/talbx/celepush/internal/pkg/model"
	"github.com/talbx/celepush/internal/pkg/service"
)

func main() {
	log.Println("celepush initialized!")
	CheckBirthdays()
	log.Println("celepush terminating!")
	os.Exit(0)
}

func CheckBirthdays() {
	log.Println("Executing birthday check")
	var appConfig model.AppConfig
	var configReader service.ConfigReader = service.ConfigReaderImpl{}
	configReader.ReadConfig(&appConfig)
	processingDelegate := createDelegate()
	processingDelegate.Delegate(appConfig)
	log.Println("Finished execution of birthday check")
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}

func createDelegate() delegate.ProcessingDelegate {
	notificationService := service.PushNotificationService{}
	messageDelegate := delegate.MessageDelegateImpl{NotificationService: notificationService}
	processingDelegate := delegate.ProcessingDelegate{MessageDelegate: messageDelegate}
	return processingDelegate
}
