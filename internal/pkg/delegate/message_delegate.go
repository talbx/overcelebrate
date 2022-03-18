package delegate

import (
	"fmt"
	"github.com/gregdel/pushover"
	"github.com/talbx/birthday-notice/internal/pkg/model"
	"github.com/talbx/birthday-notice/internal/pkg/service"
)

type MessageDelegate interface {
	Delegate(item model.BirthdayEntry, age int)
}

type MessageDelegateImpl struct {
	NotificationService service.NotificationService
}

func (delegate MessageDelegateImpl) Delegate(item model.BirthdayEntry, age int) {
	longMessage := fmt.Sprintf("%s wird heute %d ", item.Name, age)
	message := &pushover.Message{
		Message:  longMessage,
		Title:    "Geburtstag!",
		Priority: pushover.PriorityHigh,
		Sound:    pushover.SoundCosmic,
	}
	delegate.NotificationService.Notify(message)
}
