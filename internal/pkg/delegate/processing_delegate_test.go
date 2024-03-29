package delegate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/talbx/overcelebrate/internal/pkg/model"
)

type mockMessageDelegate struct {
	mock.Mock
}

func (m *mockMessageDelegate) Delegate(item model.BirthdayEntry, age int) {
	m.Called(item, age)
}

func TestShouldNotDelegateFurther(t *testing.T) {
	tomorrow := time.Now().AddDate(0, 0, 1).Format("02-01-2006")
	source := []model.BirthdayEntry{{Name: "hans", Date: tomorrow}}

	mockService := new(mockMessageDelegate)
	processor := ProcessingDelegate{mockService}
	processor.Delegate(&source)
	mockService.AssertNotCalled(t, "Delegate")
}

func TestShouldDelegateFurther(t *testing.T) {
	formatted := time.Now().Format("02-01-2006")
	hans := model.BirthdayEntry{Name: "hans", Date: formatted}
	source := []model.BirthdayEntry{hans}

	mockService := new(mockMessageDelegate)
	mockService.On("Delegate", mock.Anything, mock.Anything).Return(nil)
	processor := ProcessingDelegate{mockService}
	processor.Delegate(&source)
	mockService.Mock.AssertCalled(t, "Delegate", hans, mock.Anything)
}
