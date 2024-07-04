package counter

import (
	"strconv"
	"sync"
	"time"

	"example.com/go-demo-1/Documents/GitHub/counting-requests/src/storage"
)

type Counter struct {
	mu         sync.Mutex
	interval   time.Duration
	timestamps []time.Time
	store      storage.Storage
}

func NewCounter(interval time.Duration, store storage.Storage) *Counter {
	return &Counter{
		interval: interval,
		store:    store,
	}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	c.timestamps = append(c.timestamps, now)
	c.cleanup(now)
	c.store.Save(c.timestamps)
}

func (c *Counter) Count() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cleanup(time.Now())
	return strconv.Itoa(len(c.timestamps))
}

func (c *Counter) cleanup(now time.Time) {
	threshold := now.Add(-c.interval)
	var i int
	for i = 0; i < len(c.timestamps); i++ {
		if c.timestamps[i].After(threshold) {
			break
		}
	}
	c.timestamps = c.timestamps[i:]
}

func (c *Counter) Save() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store.Save(c.timestamps)
}

func (c *Counter) Load() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.timestamps = c.store.Load()
}
