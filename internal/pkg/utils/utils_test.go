package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSameDate(t *testing.T) {
	date1 := time.Now()
	date2 := time.Now().AddDate(1,0,0)

	isSameDate := IsSameDate(date1, date2)

	assert.True(t, isSameDate)
}

func TestNotSameDate1(t *testing.T) {
	date1 := time.Now()
	date2 := time.Now().AddDate(0,1,0)

	isSameDate := IsSameDate(date1, date2)

	assert.False(t, isSameDate)
}


func TestNotSameDate2(t *testing.T) {
	date1 := time.Now()
	date2 := time.Now().AddDate(0,0,1)

	isSameDate := IsSameDate(date1, date2)

	assert.False(t, isSameDate)
}

func TestCalculateAge(t *testing.T) {
	curr, _ := time.Parse("02-01-2006", "01-01-2020")
	bday, _ := time.Parse("02-01-2006", "01-01-2000")
	age := CalculateAge(curr.Year(), bday.Year())

	assert.Equal(t, 20, age)
}
