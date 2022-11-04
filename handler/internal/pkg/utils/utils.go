package utils

import "time"
import "go.uber.org/zap"

func CalculateAge(year int, year2 int) int {
	return year - year2
}

func IsSameDate(d1 time.Time, d2 time.Time) bool {
	sameDay := d1.Day() == d2.Day()
	sameMonth := d1.Month() == d2.Month()
	return sameDay && sameMonth
}

func L() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger) // flushes buffer, if any
	return logger.Sugar()
}

var Sugar = L()
