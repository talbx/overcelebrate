package delegate

import (
	"testing"

	"github.com/gregdel/pushover"
	"github.com/stretchr/testify/mock"
	"github.com/talbx/overcelebrate/internal/pkg/model"
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
