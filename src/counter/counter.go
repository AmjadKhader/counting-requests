package counter

import (
	"strconv"
	"sync"
	"time"

	"example.com/go-demo-1/Documents/GitHub/counting-requests/src/storage"
)

type Counter struct {
	mutex      sync.Mutex
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

func (counter *Counter) Increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	now := time.Now()
	counter.timestamps = append(counter.timestamps, now)
	counter.cleanup(now)
	counter.store.Save(counter.timestamps)
}

func (counter *Counter) Count() string {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.cleanup(time.Now())
	return strconv.Itoa(len(counter.timestamps))
}

func (counter *Counter) cleanup(now time.Time) {
	threshold := now.Add(-counter.interval)
	var i int
	for i = 0; i < len(counter.timestamps); i++ {
		if counter.timestamps[i].After(threshold) {
			break
		}
	}
	counter.timestamps = counter.timestamps[i:]
}

func (counter *Counter) Save() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.store.Save(counter.timestamps)
}

func (counter *Counter) Load() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.timestamps = counter.store.Load()
}
