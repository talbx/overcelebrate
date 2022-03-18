package delegate

import (
	"github.com/talbx/birthday-notice/internal/pkg/model"
	"github.com/talbx/birthday-notice/internal/pkg/utils"
	"log"
	"time"
)

type ProcessingDelegate struct{
	MessageDelegate MessageDelegate
}

func (delegate ProcessingDelegate) Delegate(entries *[]model.BirthdayEntry) {
	var foundSomeone = false
	for _, item := range *entries {
		now := time.Now()
		parse, err := time.Parse("02-01-2006", item.Date)
		if err != nil {
			log.Fatal(err)
		}
		if utils.IsSameDate(now, parse) {
			foundSomeone = true
			age := utils.CalculateAge(now.Year(), parse.Year())
			utils.Sugar.Infow("Found a Birthday for today!",
				"name", item.Name,
				"age", age)
			delegate.MessageDelegate.Delegate(item, age)
		}
	}
	if !foundSomeone {
		utils.Sugar.Infow("No Birthday found for today")
	}
}
