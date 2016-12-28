package mockclient

import (
	"time"

	"github.com/stffabi/watchtower/container"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) ListContainers(cf container.Filter) ([]container.Container, error) {
	args := m.Called(cf)
	return args.Get(0).([]container.Container), args.Error(1)
}

func (m *MockClient) StopContainer(c container.Container, timeout time.Duration) error {
	args := m.Called(c, timeout)
	return args.Error(0)
}

func (m *MockClient) StartContainer(c container.Container) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *MockClient) RenameContainer(c container.Container, name string) error {
	args := m.Called(c, name)
	return args.Error(0)
}

func (m *MockClient) IsContainerStale(c container.Container) (bool, error) {
	args := m.Called(c)
	return args.Bool(0), args.Error(1)
}

func (m *MockClient) RemoveImage(c container.Container) error {
	args := m.Called(c)
	return args.Error(0)
}
