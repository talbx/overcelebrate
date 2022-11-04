package overcelebrate

import (
	"log"
	"net/http"

	"github.com/talbx/overcelebrate/internal/pkg/delegate"
	"github.com/talbx/overcelebrate/internal/pkg/model"
	"github.com/talbx/overcelebrate/internal/pkg/service"
)

func CheckBirthdays() {
	log.Println("Executing birthday check")
	var appConfig model.AppConfig
	var configReader service.ConfigReader = service.ConfigReaderImpl{}
	configReader.ReadConfig(&appConfig)
	processingDelegate := createDelegate()
	processingDelegate.Delegate(appConfig)
	log.Println("Finished execution of birthday check")
}

func StartFn(w http.ResponseWriter, r *http.Request) {
	CheckBirthdays()
}

func createDelegate() delegate.ProcessingDelegate {
	notificationService := service.PushNotificationService{}
	messageDelegate := delegate.MessageDelegateImpl{NotificationService: notificationService}
	processingDelegate := delegate.ProcessingDelegate{MessageDelegate: messageDelegate}
	return processingDelegate
}
