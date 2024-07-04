package counter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	data []time.Time
}

func (m *MockStorage) Save(data []time.Time) {
	m.data = data
}

func (m *MockStorage) Load() []time.Time {
	return m.data
}

func TestCounterIncrement(t *testing.T) {
	mockStorage := &MockStorage{}
	counter := NewCounter(60*time.Second, mockStorage)

	counter.Increment()
	assert.Equal(t, "1", counter.Count())
}

func TestCounterCount(t *testing.T) {
	mockStorage := &MockStorage{}
	counter := NewCounter(60*time.Second, mockStorage)

	counter.Increment()
	time.Sleep(1 * time.Second)
	counter.Increment()
	assert.Equal(t, "2", counter.Count())
}

func TestCounterCleanup(t *testing.T) {
	mockStorage := &MockStorage{}
	counter := NewCounter(1*time.Second, mockStorage)

	counter.Increment()
	time.Sleep(2 * time.Second)
	counter.Increment()
	assert.Equal(t, "1", counter.Count())
}

func TestCounterSaveLoad(t *testing.T) {
	mockStorage := &MockStorage{}
	counter := NewCounter(60*time.Second, mockStorage)

	counter.Increment()
	counter.Save()

	newCounter := NewCounter(60*time.Second, mockStorage)
	newCounter.Load()
	assert.Equal(t, "1", newCounter.Count())
}

func TestCounterLoadEmptyFile(t *testing.T) {
	mockStorage := &MockStorage{}
	counter := NewCounter(60*time.Second, mockStorage)

	counter.Load()
	assert.Equal(t, "0", counter.Count())
}
