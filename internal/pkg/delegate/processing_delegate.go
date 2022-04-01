package delegate

import (
	"fmt"
	"github.com/talbx/celepush/internal/pkg/model"
	"github.com/talbx/celepush/internal/pkg/utils"
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
			log.Println(fmt.Sprintf("Found a Birthday for today: %s turns %d", item.Name, age))
			delegate.MessageDelegate.Delegate(item, age)
		}
	}
	if !foundSomeone {
		log.Println("No Birthday found for today")
	}
}
