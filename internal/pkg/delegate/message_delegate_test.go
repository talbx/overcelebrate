package delegate

import (
	"github.com/gregdel/pushover"
	"github.com/stretchr/testify/mock"
	"github.com/talbx/celepush/internal/pkg/model"
	"testing"
)

type mockNotificationService struct {
	mock.Mock
}

func (m *mockNotificationService) Notify(_ *pushover.Message) *pushover.Response {
	m.Called()
	return &pushover.Response{
		ID: "mock",
	}
}

func TestShouldCallNotificationService(t *testing.T) {
	item := model.BirthdayEntry{
		Name: "Hans",
		Date: "20-10-2010",
	}

	mockService := new(mockNotificationService)
	mockService.On("Notify")
	delegate := MessageDelegateImpl{mockService}
	delegate.Delegate(item, 32)
	mockService.Mock.AssertCalled(t, "Notify", mock.Anything)
}
